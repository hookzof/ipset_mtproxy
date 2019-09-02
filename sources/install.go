package main

import (
	"flag"
	"log"
	"os/exec"
)

func cmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Println(err)
		}

		return out
	}

	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Println(err)
	}

	return out
}

func main() {
	log.Println("Starting...")

	first := flag.Bool("badhosts", false, "a bool")
	second := flag.Bool("digitalocean", false, "a bool")
	third := flag.Bool("countryblock", false, "a bool")

	flag.Parse()

	log.Println("Dependency check...")
	cmd("apt -y install git unzip ipset", true)

	log.Println("Downloading and extracting...")
	cmd("cd /opt && git clone https://github.com/hookzof/ipset_mtproxy && cd ipset_mtproxy && unzip ipset.up.zip", true)
	cmd("ipset save > /etc/backup.ipset.up.rules && ipset destroy && ipset restore < /opt/ipset_mtproxy/ipset.up.rules", true)

	if *first {
		cmd("iptables -A INPUT -m set --match-set badhosts src -j DROP", true)
		log.Println("[iptables] badhosts added")
	}

	if *second {
		cmd("iptables -A INPUT -m set --match-set digitalocean src -j DROP", true)
		log.Println("[iptables] digitalocean added")
	}

	if *third {
		cmd("iptables -A INPUT -m set --match-set countryblock src -j DROP", true)
		log.Println("[iptables] countryblock added")
	}

	if *first || *second || *third {
		log.Println("[iptables] Saving rules... [/etc/rules.v4]")
		cmd("iptables-save > /etc/rules.v4", true)
	} else {
		log.Println("[iptables] Rules have not been added, check startup keys")
	}

	log.Println("Done!")
}
