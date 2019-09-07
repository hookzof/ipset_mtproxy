# ipset_mtproxy

**ipset.up.zip** содержит правила для ipset (свыше 1М записей), состоит из пяти листов:

<code>badhosts</code> (накопительный) - proxy_all.txt + IP полученные по <a href="https://t.me/unkn0wnerror/1237">скрипту</a> + IP публичных прокси (**ip:port** из открытых источников);

<code>digitalocean</code> - подсети DigitalOcean (возможное использование мощностей РКНом);

<code>countryblock</code> - подсети стран: Иран, Китай, Пакистан (потенциальные генераторы нагрузки);

<code>mikrotik</code> - IP-адреса микротиков смотрящих в Интернет (СПБ и МСК);

<code>rugov</code> - подсети госучреждений причастных к блокировкам (<a href="https://github.com/AntiZapret/antizapret/blob/master/blacklist4.txt">основная часть</a>).
<hr>

**БЫСТРАЯ УСТАНОВКА:**

```bash
curl -L -o install https://git.io/fjhCo && chmod +x install

./install -badhosts -digitalocean -rugov
```
Ошибки, которые могут встретиться:

<code>exit status 1</code> - не удалось найти строку через grep.
<hr>

**ФАЙЛЫ:**

<code>proxy_all.txt</code> - спарсенные <a href="https://lite.ip2location.com/database/px1-ip-country">прокси</a> на сентябрь 2019 года.
<hr>

**IPSET:**

Резервная копия:
```bash
ipset save > /etc/backup.ipset.up.rules
```

Сброс правил:
```bash
ipset destroy
```

Установить правила:
```bash
ipset restore < /opt/ipset_mtproxy/badhosts
ipset restore < /opt/ipset_mtproxy/digitalocean
ipset restore < /opt/ipset_mtproxy/countryblock
ipset restore < /opt/ipset_mtproxy/mikrotik
ipset restore < /opt/ipset_mtproxy/rugov

ipset save > /etc/ipset.up.rules
```

**IPTABLES:**
```bash
iptables -A INPUT -m set --match-set badhosts src -j DROP
iptables -A INPUT -m set --match-set digitalocean src -j DROP
iptables -A INPUT -m set --match-set countryblock src -j DROP
iptables -A INPUT -m set --match-set mikrotik src -j DROP
iptables -A INPUT -m set --match-set rugov src -j DROP

iptables-save > /etc/rules.v4
```
