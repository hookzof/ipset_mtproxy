package main

import (
	"flag"
	"log"
	"os/exec"
)

func cmd(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil && err.Error() != "exit status 1" && err.Error() != "exit status 2" {
		log.Println("[error]", err, "("+cmd+")")
	}

	return string(out)
}

func main() {
	log.Println("Starting!")

	/* Flag parse */

	badhosts := flag.Bool("badhosts", false, "a bool")
	mikrotik := flag.Bool("mikrotik", false, "a bool")
	countryblock := flag.Bool("countryblock", false, "a bool")
	digitalocean := flag.Bool("digitalocean", false, "a bool")
	rugov := flag.Bool("rugov", false, "a bool")

	uninstall := flag.Bool("uninstall", false, "a bool")

	flag.Parse()

	log.Println("[script] Dependency check...")
	cmd("apt -y install unzip ipset")

	if *uninstall {
		log.Println("[script] Deleting all rules and lists...")

		cmd("iptables -D INPUT -s 67.207.74.182/32 -j ACCEPT -m comment --comment \"https://test.ton.org addr\"")
		cmd("iptables -D INPUT -s 138.68.76.208/32 -j ACCEPT -m comment --comment \"ipset_mtproxy white list\"")

		cmd("iptables -D INPUT -m set --match-set badhosts src -j DROP")
		cmd("ipset destroy badhosts")

		cmd("iptables -D INPUT -m set --match-set mikrotik src -j DROP")
		cmd("ipset destroy mikrotik")

		cmd("iptables -D INPUT -m set --match-set countryblock src -j DROP")
		cmd("ipset destroy countryblock")

		cmd("iptables -D INPUT -m set --match-set digitalocean src -j DROP")
		cmd("ipset destroy digitalocean")

		cmd("iptables -D INPUT -m set --match-set rugov src -j DROP")
		cmd("ipset destroy rugov")

		log.Println("[script] Uninstall complete! Goodbye.")
		return
	}

	cmd("ipset save > /etc/backup.ipset.up.rules")
	log.Println("[backup] ipset (/etc/backup.ipset.up.rules)")
	cmd("iptables-save > /etc/backup.rules.v4")
	log.Println("[backup] iptables (/etc/backup.rules.v4)")

	/* White list */

	if cmd("iptables-save | grep \"67.207.74.182/32\"") == "" {
		cmd("iptables -I INPUT -s 67.207.74.182/32 -j ACCEPT -m comment --comment \"https://test.ton.org addr\"")
	}

	if cmd("iptables-save | grep \"138.68.76.208/32\"") == "" {
		cmd("iptables -I INPUT -s 138.68.76.208/32 -j ACCEPT -m comment --comment \"ipset_mtproxy white list\"")
	}

	log.Println("[script] Downloading and extracting...")
	cmd("cd /opt && mkdir ipset_mtproxy && cd ipset_mtproxy && " +
		"curl -L -o rules.zip https://github.com/hookzof/ipset_mtproxy/raw/master/ipset.up.zip && unzip rules.zip")

	log.Println("[iptables] Checking and deleting past iptables/ipset rules for correct operation...")

	if cmd("iptables-save | grep \"badhosts src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set badhosts src -j DROP")
		cmd("ipset destroy badhosts")
	}

	if *badhosts {
		cmd("ipset restore < /opt/ipset_mtproxy/badhosts")
		log.Println("[ipset] badhosts added")

		cmd("iptables -A INPUT -m set --match-set badhosts src -j DROP")
		log.Println("[iptables] badhosts added")
	}

	if cmd("iptables-save | grep \"mikrotik src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set mikrotik src -j DROP")
		cmd("ipset destroy mikrotik")
	}

	if *mikrotik {
		cmd("ipset restore < /opt/ipset_mtproxy/mikrotik")
		log.Println("[ipset] mikrotik added")

		cmd("iptables -A INPUT -m set --match-set mikrotik src -j DROP")
		log.Println("[iptables] mikrotik added")
	}

	if cmd("iptables-save | grep \"countryblock src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set countryblock src -j DROP")
		cmd("ipset destroy countryblock")
	}

	if *countryblock {
		cmd("ipset restore < /opt/ipset_mtproxy/countryblock")
		log.Println("[ipset] countryblock added")

		cmd("iptables -A INPUT -m set --match-set countryblock src -j DROP")
		log.Println("[iptables] countryblock added")
	}

	if cmd("iptables-save | grep \"digitalocean src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set digitalocean src -j DROP")
		cmd("ipset destroy digitalocean")
	}

	if *digitalocean {
		cmd("ipset restore < /opt/ipset_mtproxy/digitalocean")
		log.Println("[ipset] digitalocean added")

		cmd("iptables -A INPUT -m set --match-set digitalocean src -j DROP")
		log.Println("[iptables] digitalocean added")
	}

	if cmd("iptables-save | grep \"rugov src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set rugov src -j DROP")
		cmd("ipset destroy rugov")
	}

	if *rugov {
		cmd("ipset restore < /opt/ipset_mtproxy/rugov")
		log.Println("[ipset] rugov added")

		cmd("iptables -A INPUT -m set --match-set rugov src -j DROP")
		log.Println("[iptables] rugov added")
	}

	if *badhosts || *mikrotik || *countryblock || *digitalocean || *rugov {
		log.Println("[ipset] Saving rules... (/etc/ipset.up.rules)")
		cmd("ipset save > /etc/ipset.up.rules")

		log.Println("[iptables] Saving rules... (/etc/rules.v4)")
		cmd("iptables-save > /etc/rules.v4")
	} else {
		log.Println("[iptables] Rules have not been added, check startup keys")
	}

	log.Println("[script] Delete temporary files...")
	cmd("rm -r /opt/ipset_mtproxy")

	log.Println("Done!")
}
