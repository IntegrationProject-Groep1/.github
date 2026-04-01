<div align="center">

<img src="https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/banner.svg" alt="ShiftFestival Integration Platform" width="100%"/>

<br/>

[![Azure VM](https://img.shields.io/badge/Infra-Azure_VM-0089D6?logo=microsoft-azure&logoColor=white)](https://azure.microsoft.com)
[![Docker](https://img.shields.io/badge/Runtime-Docker-2496ED?logo=docker&logoColor=white)](https://docker.com)
[![RabbitMQ](https://img.shields.io/badge/Messaging-RabbitMQ-FF6600?logo=rabbitmq&logoColor=white)](#)
[![Academiejaar](https://img.shields.io/badge/Academiejaar-2025--2026-6366f1)](#)
[![CI](https://github.com/IntegrationProject-Groep1/.github/actions/workflows/generate-badges.yml/badge.svg)](https://github.com/IntegrationProject-Groep1/.github/actions/workflows/generate-badges.yml)

</div>

---

## Over Dit Project

**ShiftFestival** is het geïntegreerde evenementenplatform van de Desideriushogeschool, uitgewerkt door Integratieproject Groep 1. Het platform beheert de volledige levenscyclus van een netwerkevenement: van online inschrijving en sessieplanning tot kassabeheer, facturatie en CRM-opvolging.

Het systeem draait op een **Azure VM** met een gedistribueerde **Docker-architectuur**. Alle microservices communiceren asynchroon via **RabbitMQ** en zijn geïsoleerd in het interne Docker-netwerk `shift_net`. SSL-terminatie en reverse-proxy worden gecentraliseerd afgehandeld door de Drupal-frontend op poort 443.

> **Doel van het opleidingsonderdeel:** analyseren van klantbehoeften en deze omzetten naar een werkend, geïntegreerd prototype waarbij elke service automatisch data uitwisselt met de rest van het IT-landschap — zonder downtime en zonder handmatige synchronisatie.

---

## Architectuur

```mermaid
flowchart TD
    Internet((Internet))
    Internet -->|"HTTPS :443"| FE

    subgraph shift_net["Docker Network: shift_net"]
        FE["Frontend\nDrupal :30020\nReverse Proxy + SSL"]

        FE -->|"proxy /kassa"| K
        FE -->|"proxy /facturatie"| F
        FE -->|"proxy /planning"| P

        K["Kassa\nOdoo :8069\nPostgreSQL"]
        F["Facturatie\nFOSSBilling :80\nMySQL"]
        P["Planning\nPython :30050"]
        CRM["CRM\nSalesforce :3000"]
        MON["Monitoring\nElasticsearch :30060\nKibana :30061"]
        MQ[("RabbitMQ\nAMQP :30000\nUI :30001")]
    end

    FE  -->|publish| MQ
    K   -->|publish| MQ
    CRM -->|publish| MQ
    F   -->|publish| MQ
    P   -->|publish| MQ

    MQ -->|route| CRM
    MQ -->|route| F
    MQ -->|route| P
    MQ -->|route| MON

    HB["Heartbeat Sidecar\nelke service"]
    HB -->|"heartbeat XML 1s"| MQ
```

---

## Organisatiestructuur

```
IntegrationProject-Groep1
│
├── Project Managers (PM)
│   ├── Denis Dario
│   └── [PM 2 — naam invullen]
│
├── Team Frontend  (TL: Charles Wong)
│   ├── Jarno Janssens   — Developer / Tester
│   ├── Ilyas Fariss     — Developer / Tester
│   └── Dries Michiels   — Developer / Tester
│
├── Team Kassa  (TL: Dang Enwing)
│   ├── [Developer / Tester — naam invullen]
│   └── [Developer / Tester — naam invullen]
│
├── Team Facturatie  (TL: [naam invullen])
│   ├── [Developer / Tester — naam invullen]
│   └── [Developer / Tester — naam invullen]
│
├── Team CRM  (TL: [naam invullen])
│   ├── [Developer / Tester — naam invullen]
│   └── [Developer / Tester — naam invullen]
│
├── Team Planning  (TL: [naam invullen])
│   ├── [Developer / Tester — naam invullen]
│   └── [Developer / Tester — naam invullen]
│
└── Team Monitoring  (TL: [naam invullen])
    ├── [Developer / Tester — naam invullen]
    └── [Developer / Tester — naam invullen]
```

---

## Teams

### Team Frontend

![Frontend](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/frontend.svg)
[![PHP](https://img.shields.io/badge/PHP-8.2-777BB4?logo=php&logoColor=white)](#)
[![Drupal](https://img.shields.io/badge/CMS-Drupal_10-0678BE?logo=drupal&logoColor=white)](#)
[![MariaDB](https://img.shields.io/badge/DB-MariaDB-003545?logo=mariadb&logoColor=white)](#)

**Repository:** [IP-groep1-frontend](https://github.com/IntegrationProject-Groep1/IP-groep1-frontend)  
**Team Lead:** Charles Wong · **Devs:** Jarno Janssens, Ilyas Fariss, Dries Michiels

De publieke voordeur van het platform. Drupal 10 verzorgt de bezoekerswebsite waar deelnemers zich kunnen inschrijven voor sessies. De container fungeert tevens als SSL-terminatie en reverse proxy: alle HTTPS-verkeer op poort 443 wordt doorgestuurd naar de juiste interne dienst.

| Poort | Rol |
|-------|-----|
| `30020` | HTTP (intern / lokaal) |
| `443` | HTTPS (via Azure-proxy) |
| `30000` | RabbitMQ AMQP (outbound) |

**Berichtenstromen (RabbitMQ):**
- `frontend.user.registered` naar CRM (nieuwe inschrijving)
- `frontend.user.checkin` naar Kassa / CRM (check-in aan de kassa)
- `frontend.session.update` ontvangen van Planning

---

### Team Kassa

![Kassa](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/kassa.svg)
[![Python](https://img.shields.io/badge/Python-3.12-3776AB?logo=python&logoColor=white)](#)
[![Odoo](https://img.shields.io/badge/POS-Odoo_17-714B67?logo=odoo&logoColor=white)](#)
[![PostgreSQL](https://img.shields.io/badge/DB-PostgreSQL_15-4169E1?logo=postgresql&logoColor=white)](#)

**Repository:** [Kassa](https://github.com/IntegrationProject-Groep1/Kassa)  
**Team Lead:** Dang Enwing · **Devs:** [naam invullen], [naam invullen]

Het fysieke kassasysteem tijdens het festival. Odoo 17 Point-of-Sale verwerkt consumptiebestellingen en registreert betalingen. Een Python-integratiecontainer communiceert via de Odoo XML-RPC API om verkoopdata door te sturen naar de rest van het platform — zonder dat er code in Odoo zelf geschreven wordt.

| Poort | Rol |
|-------|-----|
| `8069` | Odoo Web UI + XML-RPC API |
| `5432` | PostgreSQL (intern) |

**Berichtenstromen (RabbitMQ):**
- naar CRM: consumptie gekoppeld aan klant
- naar Facturatie: betaling verwerkt (trigger factuurstatus)

---

### Team Facturatie

![Facturatie](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/facturatie.svg)
[![Python](https://img.shields.io/badge/Python-3.12-3776AB?logo=python&logoColor=white)](#)
[![FOSSBilling](https://img.shields.io/badge/Billing-FOSSBilling-E63946)](#)
[![MySQL](https://img.shields.io/badge/DB-MySQL_8-4479A1?logo=mysql&logoColor=white)](#)

**Repository:** [Facturatie](https://github.com/IntegrationProject-Groep1/Facturatie)  
**Team Lead:** [naam invullen] · **Devs:** [naam invullen], [naam invullen]

De financiële backbone van het platform. FOSSBilling maakt automatisch facturen aan voor bedrijven op basis van inschrijvingen en consumptiedata van de Kassa. Bij annuleringen worden creditnota's aangemaakt. Na afloop van het event worden alle openstaande facturen automatisch gesloten en doorgestuurd naar het mailingteam.

| Poort | Rol |
|-------|-----|
| `80` | FOSSBilling Web UI (intern) |
| `3306` | MySQL (intern) |

**Berichtenstromen (RabbitMQ):**
- van CRM: inschrijving / consumptie / betaling / annulering
- naar Mailing: factuur klaar voor verzending
- naar DLQ: foutieve berichten (Dead Letter Queue)

---

### Team CRM

![CRM](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/crm.svg)
[![JavaScript](https://img.shields.io/badge/Node.js-20-339933?logo=node.js&logoColor=white)](#)
[![Salesforce](https://img.shields.io/badge/CRM-Salesforce-00A1E0?logo=salesforce&logoColor=white)](#)

**Repository:** [CRM](https://github.com/IntegrationProject-Groep1/CRM)  
**Team Lead:** [naam invullen] · **Devs:** [naam invullen], [naam invullen]

Het geheugen van het platform. Salesforce slaat alle contacten (personen en bedrijven) op en houdt activiteiten bij. Een Node.js-consumer luistert op RabbitMQ en synchroniseert binnenkomende data automatisch naar Salesforce. Na het event kunnen de contactlijsten gebruikt worden om zakelijke relaties verder uit te bouwen.

| Poort | Rol |
|-------|-----|
| `3000` | CRM Receiver API (intern) |
| `5672` | RabbitMQ AMQP (intern) |
| `15672` | RabbitMQ Management UI |

**Berichtenstromen (RabbitMQ):**
- van Frontend: nieuwe inschrijving
- van Kassa: betaling / consumptie
- naar Facturatie: klantgegevens doorgeven
- naar Mailing: contactlijsten aanleveren

---

### Team Planning

![Planning](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/planning.svg)
[![Python](https://img.shields.io/badge/Python-3.12-3776AB?logo=python&logoColor=white)](#)
[![Microsoft Graph](https://img.shields.io/badge/Graph_API-Microsoft-0078D4?logo=microsoft&logoColor=white)](#)

**Repository:** [Planning](https://github.com/IntegrationProject-Groep1/Planning)  
**Team Lead:** [naam invullen] · **Devs:** [naam invullen], [naam invullen]

Beheert de sessie-agenda van het event. De service ontvangt `calendar.invite`-berichten en publiceert `session.created`-events naar andere teams. Via de Microsoft Graph API (OAuth 2.0) kunnen sessies rechtstreeks als event in de Outlook-kalender van een deelnemer worden aangemaakt.

| Poort | Rol |
|-------|-----|
| `30050` | Health endpoint + API |
| `30000` | RabbitMQ AMQP (outbound) |

**Berichtenstromen (RabbitMQ):**
- van Frontend: `calendar.invite`
- naar alle teams: `planning.session.created`

---

### Team Monitoring

![Monitoring](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/monitoring.svg)
[![Elasticsearch](https://img.shields.io/badge/Search-Elasticsearch_8-00BFB3?logo=elasticsearch&logoColor=white)](#)
[![Kibana](https://img.shields.io/badge/Dashboard-Kibana_8-E8478B?logo=kibana&logoColor=white)](#)
[![Logstash](https://img.shields.io/badge/Pipeline-Logstash_8-FEC514?logo=logstash&logoColor=white)](#)

**Repository:** [Monitoring](https://github.com/IntegrationProject-Groep1/Monitoring)  
**Team Lead:** [naam invullen] · **Devs:** [naam invullen], [naam invullen]

De controlroom van het platform. De ELK-stack (Elasticsearch + Logstash + Kibana) ontvangt elke seconde een XML-heartbeat van elke service via RabbitMQ. Logstash parsed en indexeert de data; Kibana visualiseert de uptime van alle teams in realtime. Berichten met ongeldige XML of onbekende systemen gaan naar een quarantine-index.

| Poort | Rol |
|-------|-----|
| `30060` | Elasticsearch REST API |
| `30061` | Kibana Dashboard UI |

**Gemonitorde systemen:** `frontend` · `kassa` · `facturatie` · `crm` · `planning` · `monitoring`

---

### Heartbeat Sidecar

![Heartbeat](https://raw.githubusercontent.com/IntegrationProject-Groep1/.github/main/profile/badges/heartbeat.svg)
[![Python](https://img.shields.io/badge/Python-3.12-3776AB?logo=python&logoColor=white)](#)
[![GHCR](https://img.shields.io/badge/Image-GHCR-181717?logo=github&logoColor=white)](https://github.com/IntegrationProject-Groep1/heartbeat/pkgs/container/heartbeat)

**Repository:** [heartbeat](https://github.com/IntegrationProject-Groep1/heartbeat)

Een gedeelde sidecar-container die elk team toevoegt aan zijn eigen `docker-compose.yml`. De sidecar controleert elke seconde of de opgegeven containers bereikbaar zijn via TCP en stuurt een heartbeat-XML naar RabbitMQ. Team Infrastructuur beheert de deployments via Watchtower — nieuwe versies worden automatisch uitgerold.

```yaml
# Toevoegen aan jullie docker-compose.yml:
sidecar:
  image: ghcr.io/integrationproject-groep1/heartbeat:latest
  environment:
    - SYSTEM_NAME=jullie-systeem-naam   # bijv. kassa / crm / planning
    - TARGETS=container-naam:poort
    - RABBITMQ_HOST=rabbitmq_broker
    - RABBITMQ_USER=<username>
    - RABBITMQ_PASS=<password>
```

---

## Poortenoverzicht

| Team | Service | Host-poort | Gebruik |
|------|---------|-----------|---------|
| Frontend | Drupal (HTTP) | `30020` | Web UI & registratie |
| Frontend | Apache (HTTPS proxy) | `443` | SSL-terminatie & reverse proxy |
| Kassa | Odoo Web + XML-RPC | `8069` | Point of Sale UI & API |
| Facturatie | FOSSBilling | `80` (intern) | Factuur-beheer UI |
| CRM | Node.js Receiver | `3000` | CRM-integratie API |
| Planning | Python Health/API | `30050` | Health endpoint & planning API |
| Monitoring | Elasticsearch | `30060` | REST API (intern) |
| Monitoring | Kibana | `30061` | Dashboard UI |
| Infra | RabbitMQ AMQP | `30000` | Berichtenwachtrij (AMQP) |
| Infra | RabbitMQ Management | `30001` | Beheer UI |

---

## Repositories

| Repo | Technologie | Omschrijving |
|------|-------------|--------------|
| [IP-groep1-frontend](https://github.com/IntegrationProject-Groep1/IP-groep1-frontend) | ![PHP](https://img.shields.io/badge/-PHP-777BB4?logo=php&logoColor=white) | Drupal 10 — publieke website, inschrijvingen, SSL proxy |
| [Kassa](https://github.com/IntegrationProject-Groep1/Kassa) | ![Python](https://img.shields.io/badge/-Python-3776AB?logo=python&logoColor=white) | Odoo 17 Point of Sale + Python integratie |
| [Facturatie](https://github.com/IntegrationProject-Groep1/Facturatie) | ![Python](https://img.shields.io/badge/-Python-3776AB?logo=python&logoColor=white) | FOSSBilling — facturen, creditnota's, betalingen |
| [CRM](https://github.com/IntegrationProject-Groep1/CRM) | ![JS](https://img.shields.io/badge/-JavaScript-F7DF1E?logo=javascript&logoColor=black) | Salesforce CRM — contactbeheer & mailinglijsten |
| [Planning](https://github.com/IntegrationProject-Groep1/Planning) | ![Python](https://img.shields.io/badge/-Python-3776AB?logo=python&logoColor=white) | Sessieplanning + Microsoft Graph API (Outlook) |
| [Monitoring](https://github.com/IntegrationProject-Groep1/Monitoring) | ![Python](https://img.shields.io/badge/-Python-3776AB?logo=python&logoColor=white) | ELK Stack — realtime uptime dashboard |
| [heartbeat](https://github.com/IntegrationProject-Groep1/heartbeat) | ![Python](https://img.shields.io/badge/-Python-3776AB?logo=python&logoColor=white) | Gedeelde heartbeat sidecar voor alle teams |
| [.github](https://github.com/IntegrationProject-Groep1/.github) | ![Go](https://img.shields.io/badge/-Go-00ADD8?logo=go&logoColor=white) | Org-profiel, badge-generator & workflows |

---

<div align="center">

Badges worden automatisch gegenereerd via GitHub Actions + Go — zie [`badge-generator/`](https://github.com/IntegrationProject-Groep1/.github/tree/main/badge-generator)

[![Generate Badges](https://github.com/IntegrationProject-Groep1/.github/actions/workflows/generate-badges.yml/badge.svg)](https://github.com/IntegrationProject-Groep1/.github/actions/workflows/generate-badges.yml)

</div>
