# ipset_mtproxy

**ipset.up.zip** содержит правила для ipset (свыше 1М записей), включает в себя 2 листа:

<code>badhosts</code> - proxy_all.txt + IP полученные по скрипту - https://t.me/c/1301206189/5916 (накопляется по мере обновления листов firehol)<br>
<code>countryblock</code> - госы (основная часть - https://github.com/AntiZapret/antizapret/blob/master/blacklist4.txt), из стран: Иран, Китай, Пакистан (потенциальные генераторы нагрузки)
<hr>

Спарсенные прокси (https://lite.ip2location.com/database/px1-ip-country) на август 2019 года:

<code>proxy_ru.txt</code> - RU список<br>
<code>proxy_all.txt</code> - полный список
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
iptables -A INPUT -m set --match-set countryblock src -j DROP
iptables -A INPUT -m set --match-set badhosts src -j DROP

iptables-save > /etc/rules.v4
```
