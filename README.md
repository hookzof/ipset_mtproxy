# ipset_mtproxy

⚠️ При использовании этих правил возрастает потребление оперативной памяти от 500 МБ!

**ipset.up.zip** содержит правила для ipset (свыше 1М записей), состоит из 3-х листов:

<code>badhosts</code> (накопительный) - proxy_all.txt + IP полученные по <a href="https://t.me/unkn0wnerror/1237">скрипту</a> + IP публичных прокси (полученные по маске **IP:PORT** из открытых источников);

<code>digitalocean</code> - подсети DigitalOcean (возможное использование мощностей РКНом);

<code>countryblock</code>:
- подсети госучреждений причастных к блокировкам (<a href="https://github.com/AntiZapret/antizapret/blob/master/blacklist4.txt">основная часть</a>);
- подсети стран: Иран, Китай, Пакистан (потенциальные генераторы нагрузки).
<hr>

**БЫСТРАЯ УСТАНОВКА:**

```bash
wget https://github.com/hookzof/ipset_mtproxy/raw/master/install && chmod +x install

./install -badhosts -digitalocean
```

P.S. Не для многократного использования (ломается iptables и пр.)
<hr>

**ФАЙЛЫ:**

<code>proxy_all.txt</code> - спарсенные <a href="https://lite.ip2location.com/database/px1-ip-country">прокси</a> на сентябрь 2019 года;

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
```

Установить правила:
```bash
ipset restore < /etc/ipset.up.rules
```

**IPTABLES:**
```bash
iptables -A INPUT -m set --match-set badhosts src -j DROP
iptables -A INPUT -m set --match-set digitalocean src -j DROP
iptables -A INPUT -m set --match-set countryblock src -j DROP

iptables-save > /etc/rules.v4
```
