/*
    YARA Rules for Offensive Agent Detection
    
    Purpose: Educational - Teach defenders how to detect this tool
    
    These rules are provided to help:
    - Blue teams detect this tool in their environments
    - SOC analysts create detection signatures
    - Students learn YARA rule writing
    - Security researchers identify the tool
    
    IMPORTANT: These rules do NOT absolve the author of legal responsibility.
    They are provided as an educational resource for defenders.
*/

rule OffensiveAgent_Registry_Persistence_Strings
{
    meta:
        description = "Detects Go-based offensive agent with registry persistence capability"
        author = "Educational Security Research"
        date = "2025-02-18"
        severity = "medium"
        reference = "https://github.com/iojabal/offensive-agent"
        purpose = "Educational detection for authorized testing environments"
        mitre_attack = "T1547.001"
        
    strings:
        // Registry persistence indicators
        $reg1 = "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run" ascii wide
        $reg2 = "HKLM\\Software\\Microsoft\\Windows\\CurrentVersion\\Run" ascii wide
        $reg3 = "SysBackdoor" ascii wide
        
        // Go-specific strings
        $go1 = "go.buildid" ascii
        $go2 = "runtime.main" ascii
        
        // Command functionality
        $cmd1 = "persistence enable" ascii
        $cmd2 = "persistence disable" ascii
        $cmd3 = "persistence status" ascii
        
        // File paths from your project
        $path1 = "commands/persistence" ascii
        $path2 = "RegPersist.go" ascii
        $path3 = "windows.go" ascii
        
    condition:
        uint16(0) == 0x5A4D and  // PE file
        (
            (2 of ($reg*)) and 
            (1 of ($go*)) and 
            (1 of ($cmd*))
        ) or
        (3 of ($path*))
}

rule OffensiveAgent_Command_Structure
{
    meta:
        description = "Detects command structure of educational offensive agent"
        author = "Educational Security Research"
        date = "2025-02-18"
        severity = "low"
        purpose = "Blue team training and detection practice"
        
    strings:
        // Command dispatcher patterns
        $cmd1 = "ReconCommand" ascii
        $cmd2 = "PersistenceCommand" ascii
        $cmd3 = "InfoCommand" ascii
        $cmd4 = "HelpCommand" ascii
        
        // Package indicators
        $pkg1 = "package persistence" ascii
        $pkg2 = "package recon" ascii
        $pkg3 = "package commands" ascii
        
        // Function signatures
        $func1 = "func IsElevated()" ascii
        $func2 = "func RegPersist()" ascii
        
    condition:
        (3 of ($cmd*)) or
        (2 of ($pkg*) and 1 of ($func*))
}

rule OffensiveAgent_Network_Indicators
{
    meta:
        description = "Network behavior patterns of educational agent"
        author = "Educational Security Research"
        date = "2025-02-18"
        severity = "medium"
        purpose = "Network-based detection for blue team training"
        
    strings:
        // Transport layer indicators
        $net1 = "transport.tcp" ascii
        $net2 = "net.Dial" ascii
        $net3 = "bufio.NewReader" ascii
        
        // Session management
        $sess1 = "Session control" ascii
        $sess2 = "command routing" ascii
        
    condition:
        (2 of ($net*)) and (1 of ($sess*))
}

rule OffensiveAgent_Complete_Signature
{
    meta:
        description = "Comprehensive signature for the educational offensive agent"
        author = "Educational Security Research"
        date = "2025-02-18"
        severity = "high"
        purpose = "Complete detection profile for SOC operations"
        reference = "https://github.com/iojabal/offensive-agent"
        mitre_attack = "T1547.001, TA0011"
        confidence = "high"
        
    strings:
        // Unique project identifiers
        $id1 = "offensive-agent" ascii wide
        $id2 = "nombredetuapp" ascii  // Your package path
        
        // Registry persistence
        $pers1 = "Registry key created for persistence" ascii
        $pers2 = "Persistence Enabled:" ascii
        
        // Help text indicators
        $help1 = "persistence <enable|disable|status>" ascii
        $help2 = "Usage: recon" ascii
        
        // Error messages
        $err1 = "Persistence command executed" ascii
        $err2 = "Unknown action. Use 'enable', 'disable', or 'status'" ascii
        
    condition:
        uint16(0) == 0x5A4D and  // PE file
        (
            (1 of ($id*)) or
            (2 of ($pers*)) or
            (2 of ($help*))
        )
}

/*
    Detection Recommendations for Blue Teams:
    
    1. Deploy these rules in:
       - Endpoint detection tools
       - SIEM platforms
       - Threat hunting workflows
       
    2. Combine with behavioral detection:
       - Registry modifications at startup
       - Unusual process executions
       - Network connections to non-standard ports
       
    3. Response procedures:
       - Isolate affected system
       - Collect forensic evidence
       - Investigate authorization
       - Review security policies
       
    4. Prevention strategies:
       - Application whitelisting
       - Registry monitoring
       - User education
       - Least privilege principle
*/
