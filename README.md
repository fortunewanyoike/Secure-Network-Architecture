# Secure Network Architecture: Segmentation & Access Control Implementation

## 🛡️ Project Overview
This project demonstrates the design and implementation of a **Secure Network Architecture** for a sensitive healthcare environment. By transitioning from a flat network to a segmented topology, this project enforces the **Principle of Least Privilege (PoLP)** and mitigates the risk of lateral movement. The implementation features a **VyOS router** acting as a stateful firewall to isolate a database server from unauthorized departmental subnets.

### 🎯 Objectives
- Design a segmented network topology to isolate sensitive healthcare databases.
- Implement **Layer 4 stateful firewall rules** to control traffic between Finance, HR, and Database subnets.
- Configure persistent static routing and **Netplan** settings for multi-subnet communication.
- Perform real-time traffic analysis using **tcpdump** to validate security policy enforcement.

---

## 🛠️ Technical Implementation
The architecture was built and validated in a virtualized lab environment using enterprise-grade networking tools.

### Key Components:
- **VyOS Router**: Serves as the central security gateway and firewall enforcement point.
- **Network Segmentation**: 
  - **Finance Subnet** (192.168.1.0/24): Authorized for database access.
  - **HR Subnet** (192.168.2.0/24): Restricted from database access.
  - **Database Subnet** (192.168.3.0/24): Isolated "Walled Garden" for sensitive data.
- **Access Control Lists (ACLs)**:
  - **Permit**: TCP port 3306 (MySQL) from Finance to Database.
  - **Drop & Log**: All traffic from HR to Database.

---

## 🔍 Security Analysis & Validation
The effectiveness of the segmentation was verified through rigorous testing and packet analysis.

### 1. Firewall Enforcement
- **Success**: Finance-VM successfully established a TCP three-way handshake with the MySQL service.
- **Blocked**: HR-VM connection attempts resulted in timeouts, with the VyOS router silently dropping packets as per the "Drop" policy.

### 2. Traffic Inspection
- Utilized `tcpdump` on the Database interface to confirm that only authorized source IPs were reaching the server.
- Verified that return traffic was correctly routed back through the VyOS gateway using persistent static routes.

---

## 📈 Impact & Best Practices
- **Blast Radius Reduction**: By isolating the HR subnet, a compromise of an HR workstation cannot propagate to the critical database.
- **CIS Benchmark Alignment**: Follows industry standards for network hardening and centralized logging.
- **Ransomware Mitigation**: Limits the ability of automated malware to scan and infect internal database servers.

---

## 📝 Future Enhancements
- **Intrusion Detection**: Integrating an **IDS (Snort/Suricata)** to inspect allowed traffic for application-layer attacks like SQL Injection.
- **Centralized Logging**: Exporting VyOS firewall logs to a **SIEM** (e.g., ELK Stack) for real-time alerting on unauthorized access attempts.

---
**Author**: Fortune Migwi Wanyoike  
**Role**: Junior Cybersecurity Engineer  
**Technologies**: VyOS, Linux Networking (Netplan), tcpdump, MySQL, Firewall/ACLs
