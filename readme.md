# Clockr ⏰

Clockr is a lightweight desktop time-punching app made for clocking hours at work.

I wrote this with Go, it is my second ever Go project.

This is a personal project tailored for my job, this probably won't be perfect for everybody's use case.

---

## Features

- **Clock In / Clock Out** – start and stop your work sessions 
- **Break Tracking** – start and end breaks during an active session 
- **Live Status & Summary** – always know how long you’ve been working or resting  
- **Export Hours** – generate logs of your sessions for spreadsheets, payroll, or existential reflection  
- **Adjust Timesheet** – manually adjust your clock in/out and break start/end times in case you edit your session on accident

---

## Installation

1. Clone the repo:  
```bash
git clone https://github.com/kikiluvv/clockr.git
cd clockr
```

2. Make sure you have Go installed (1.20+ recommended):
```bash
go version
```

3. Build the app:
```bash
go build -o clockr ./gui
```

4. Run it:
```bash
./clockr
```

---

## Usage

### Commands

- **gui**: Open GUI
- **in**: Clock in for work
- **out**: Clock out from work
- **break start**: Start a break
- **break end**: End the current break
- **status**: Show current session status
- **summary**: Show quick summary of hours worked (day/week/pay period)
- **export**: Export hours for the current pay period
- **adjust <date> in|out <HH:MM>**: Adjust clock in/out time for a session
- **config show**: View the current config
- **config set <field> <value> <pay_period/start/end>**: Set a config field

### GUI
- GUI incomplete - Coming Soon

---
