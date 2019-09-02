# ipset_mtproxy

**ipset.up.zip** содержит правила для ipset (свыше 1М записей), состоит из 3-х листов:

<code>badhosts</code> (накопительный) - proxy_all.txt, IP полученные по скрипту - https://t.me/c/1301206189/5916, IP публичных прокси, полученные по маске <code>IP:PORT</code> из открытых источников;<br>

<code>digitalocean</code> - подсети DigitalOcean (возможное использование мощностей РКНом); <br/>

<code>countryblock</code>:
- подсети госучреждений причастных к блокировкам (основная часть - https://github.com/AntiZapret/antizapret/blob/master/blacklist4.txt); <br/>
- подсети стран: Иран, Китай, Пакистан (потенциальные генераторы нагрузки).
<hr>

<code>proxy_all.txt</code> - спарсенные прокси (https://lite.ip2location.com/database/px1-ip-country) на сентябрь 2019 года;<br>
<code>mikrotik_test.txt</code> - СПБ и МСК микротики смотрящие по 1723 и 2000 портам в Интернет.
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
