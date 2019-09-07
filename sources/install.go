package main

import (
	"flag"
	"log"
	"os/exec"
)

func cmd(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Println(err)
	}

	return string(out)
}

func main() {
	log.Println("Starting!")

	first := flag.Bool("badhosts", false, "a bool")
	second := flag.Bool("digitalocean", false, "a bool")
	third := flag.Bool("countryblock", false, "a bool")
	fourth := flag.Bool("rugov", false, "a bool")
	fifth := flag.Bool("mikrotik", false, "a bool")

	flag.Parse()

	log.Println("[system] Dependency check...")
	cmd("apt -y install git unzip ipset")

	cmd("ipset save > /etc/backup.ipset.up.rules")
	log.Println("[backup] ipset (/etc/backup.ipset.up.rules)")
	cmd("iptables-save > /etc/backup.rules.v4")
	log.Println("[backup] iptables (/etc/backup.rules.v4)")

	log.Println("[system] Downloading and extracting...")
	cmd("cd /opt && git clone https://github.com/hookzof/ipset_mtproxy && cd ipset_mtproxy && unzip ipset.up.zip")

	log.Println("[iptables] Checking and deleting past iptables rules for correct operation...")

	if cmd("iptables-save | grep \"badhosts src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set badhosts src -j DROP")
	}

	if cmd("iptables-save | grep \"digitalocean src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set digitalocean src -j DROP")
	}

	if cmd("iptables-save | grep \"countryblock src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set countryblock src -j DROP")
	}

	if cmd("iptables-save | grep \"rugov src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set rugov src -j DROP")
	}

	if cmd("iptables-save | grep \"mikrotik src\"") != "" {
		cmd("iptables -D INPUT -m set --match-set mikrotik src -j DROP")
	}

	log.Println("[ipset] An attempt to destroy ipset and restore rules...")
	cmd("ipset destroy")

	if *first {
		cmd("ipset restore < /opt/ipset_mtproxy/badhosts")
		log.Println("[ipset] badhosts added")
		cmd("iptables -A INPUT -m set --match-set badhosts src -j DROP")
		log.Println("[iptables] badhosts added")
	}

	if *second {
		cmd("ipset restore < /opt/ipset_mtproxy/digitalocean")
		log.Println("[ipset] digitalocean added")
		cmd("iptables -A INPUT -m set --match-set digitalocean src -j DROP")
		log.Println("[iptables] digitalocean added")
	}

	if *third {
		cmd("ipset restore < /opt/ipset_mtproxy/countryblock")
		log.Println("[ipset] countryblock added")
		cmd("iptables -A INPUT -m set --match-set countryblock src -j DROP")
		log.Println("[iptables] countryblock added")
	}

	if *fourth {
		cmd("ipset restore < /opt/ipset_mtproxy/rugov")
		log.Println("[ipset] rugov added")
		cmd("iptables -A INPUT -m set --match-set rugov src -j DROP")
		log.Println("[iptables] rugov added")
	}

	if *fifth {
		cmd("ipset restore < /opt/ipset_mtproxy/mikrotik")
		log.Println("[ipset] mikrotik added")
		cmd("iptables -A INPUT -m set --match-set mikrotik src -j DROP")
		log.Println("[iptables] mikrotik added")
	}

	if *first || *second || *third || *fourth || *fifth {
		log.Println("[ipset] Saving rules... (/etc/ipset.up.rules)")
		cmd("ipset save > /etc/ipset.up.rules")
		log.Println("[iptables] Saving rules... (/etc/rules.v4)")
		cmd("iptables-save > /etc/rules.v4")
	} else {
		log.Println("[iptables] Rules have not been added, check startup keys")
	}

	log.Println("[system] Delete temporary files...")
	cmd("rm -r /opt/ipset_mtproxy")

	log.Println("Done!")
}
