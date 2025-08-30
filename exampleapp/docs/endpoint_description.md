Tento endpoint slouží k získání informací o konkrétním uživateli na základě jeho jedinečného ID.

### Použití

Odešlete GET požadavek na `/uzivatele/{id}` kde `{id}` je číselný identifikátor uživatele.

### Návratové hodnoty

- **200 OK**: Uživatel byl nalezen a jeho data jsou vrácena
- **404 Not Found**: Uživatel s daným ID neexistuje
- **500 Internal Server Error**: Chyba serveru

### Bezpečnost

Tento endpoint vyžaduje platný autentifikační token v Authorization hlavičce.