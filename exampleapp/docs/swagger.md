<details>
  <summary><b><font size="5">Základní aplikační logika</font></b></summary>
  <p>Toto REST API k elektronické komunikaci pro dodavatele s architekturou REST a formátem dat JSON s kódováním UTF8. Autentizace probíhá na základě certifikátu.</p>
</details>

<details>
  <summary><b><font size="5">Operace webové služby</font></b></summary>

### Zaevidování zprávy zadavateli
Pro zaevidování zprávy se používá endpoint **POST /servisni/zprava-zadavali**. Očekává požadavek typu `multipart/form-data` se dvěma částmi: `zpravaJson` (JSON s daty zprávy) a volitelnými `prilohy` (soubory s maximální celkovou velikostí 1 GB).

### Zaevidování přečtení zprávy zadavali
Endpoint **POST /servisni/zprava-zadavali/precteni** slouží k zaevidování přečtení zprávy. Přečtení lze zaevidovat pouze u zprávy, která byla zaevidována stejným elektronickým nástrojem.

### Detail zprávy zadavateli
Pro načtení detailu zprávy je k dispozici endpoint **GET /servisni/zprava-zadavali/{id}**, kde `{id}` je UUID zprávy.

### Smazání zprávy zadavateli
Endpoint **DELETE /servisni/zprava-zadavali/{id}** umožňuje smazání zprávy. Smazat lze pouze zprávy, které pocházejí ze stejného elektronického nástroje, který je evidoval.
</details>

---

<b><font size="5">Seznam změn</font></b>
### Aktuální verze

<details>
  <summary>1.0.0</summary>

#### Inicializační verze
- Implementace základních operací pro správu zpráv.
- Podpora zaevidování zprávy s přílohami.
- Možnost zaznamenat přečtení zprávy a získat její detail.
- Funkce pro smazání zprávy s omezením na nástroj, který ji vytvořil.
</details>
