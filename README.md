# ipset_mtproxy

**ipset.up.zip** содержит правила для ipset (свыше 1М записей), состоит из двух листов:

<code>badhosts</code> (накопительный) - proxy_all.txt, IP полученные по скрипту - https://t.me/c/1301206189/5916, IP публичных прокси, полученные по маске <code>IP:PORT</code> из открытых источников.<br>

<code>digitalocean</code> - подсети DigitalOcean (возможное использование мощностей РКНом). <br/>
<code>countryblock</code> 
- подсети государственных учереждений причастных к блокировкам (основная часть - https://github.com/AntiZapret/antizapret/blob/master/blacklist4.txt), <br/>
- подсети из из стран: Иран, Китай, Пакистан (потенциальные генераторы нагрузки), 
<hr>

<code>proxy_all.txt</code> - Спарсенные прокси (https://lite.ip2location.com/database/px1-ip-country) на август 2019 года:
<hr>

**IPSET:**

Резервная копия:
```bash
ipset save > /etc/backup.ipset.up.rules
```

Сброс правил:
```bash
ipset destroy

# не всегда работает, самый гарантированный способ сбрасывать через перезагрузку
```

Установить правила:
```bash
ipset restore < /etc/ipset.up.rules

# путь может быть другой
```

**IPTABLES:**
```bash
iptables -A INPUT -m set --match-set badhosts src -j DROP
iptables -A INPUT -m set --match-set digitalocean src -j DROP
iptables -A INPUT -m set --match-set countryblock src -j DROP

iptables-save > /etc/rules.v4
```
