# OS Concepts

<br />

## List of Contents:
### 1. [Roadmap.sh - Networking](#content-1)
### 2. [What is the OSI Model?](#content-2)
### 3. [OSI Model](#content-3)
### 4. [What Is Computer Networking?](#content-4)
### 5. [POSIX](#content-5)
### 6. [What is the meaning of "POSIX"?](#content-6)
### 7. [What is POSIX? Why Does it Matter to Linux/UNIX Users?](#content-7)
### 8. [What is a Network Socket?](#content-8)
### 9. [What is a Socket](#content-9)
### 10. [Socket Programming in Computer Network](#content-10)
### 11. [A quick introduction to processes in Computer Science](#content-11)
### 12. [Operating System - Processes](#content-12)
### 13. [I/O Hardware](#content-13)
### 14. [Virtualization](#content-14)
### 15. [What is virtualization?](#content-15)
### 16. [Containers vs. virtual machines](#content-16)
### 17. [File System](#content-17)
### 18. [Demystifying memory management in modern programming languages](#content-18)
### 19. [Operating System: Threads and Concurrency](#content-19)

<br />

---

## Contents:

## [Roadmap.sh - Networking](https://roadmap.sh/devops) <span id="content-1"></span>
- Computer networking refers to interconnected computing devices that can exchange data and share resources with each other.
- These networked devices use a system of rules, called communications protocols.

## [What is the OSI Model?](https://www.cloudflare.com/learning/ddos/glossary/open-systems-interconnection-model-osi/) <span id="content-2"></span>

### What is the OSI Model?
- OSI: Open Systems Interconnection
- The OSI provides a standard for different computer systems to be able to communicate with each other.
- The OSI Model can be seen as a universal language for computer networking.
- It is based on the concept of splitting up a communication system into seven abstract layers, each one stacked upon the last.
  ![OSI Model 7 Layers](https://cf-assets.www.cloudflare.com/slt3lc6tev37/6ZH2Etm3LlFHTgmkjLmkxp/59ff240fb3ebdc7794ffaa6e1d69b7c2/osi_model_7_layers.png)
- DDoS attacks target specific layers of a network connection; application layer attacks target layer 7 and protocol layer attacks target layers 3 and 4.
- Let's create a donkey bridge: A Presentation Suitable for Teen, Never Does Pay

### What are the 7 layers of the OSI Model?
- The application layer (7)
  ![Application Layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/2rcDKpr4WLqoyAZ7GDKkyJ/7cab96402de7ac5465b86e617da3da4e/osi_model_application_layer_7.png)
  - This is the only layer that directly interacts with data from the user.
  - Software applications like web browsers and email clients rely on the application layer to initiate communications.
  - But it should be made clear that client software applications are not part of the application layer; rather the application layer is responsible for the protocols and data manipulation that the software relies on to present meaningful data to the user.
  - Application layer protocols include HTTP as well as SMTP (Simple Mail Transfer Protocol is one of the protocols that enables email communications).
- The presentation layer (6)
  ![Presentation Layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/19L86neKKT8srUkOSe4rf7/ff4c91c94a1790651df7b48433913f59/osi_model_presentation_layer_6.png)
  - Layer 6 makes the data presentable for applications to consume.
  - The presentation layer is responsible for translation, encryption, and compression of data.
  - Two communicating devices communicating may be using different encoding methods, so layer 6 is responsible for translating incoming data into a syntax that the application layer of the receiving device can understand.
  - If the devices are communicating over an encrypted connection, layer 6 is responsible for adding the encryption on the sender’s end as well as decoding the encryption on the receiver's end so that it can present the application layer with unencrypted, readable data.
  - Finally the presentation layer is also responsible for compressing data it receives from the application layer before delivering it to layer 5. This helps improve the speed and efficiency of communication by minimizing the amount of data that will be transferred.
- The session layer (5)
  ![Session Layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/29mRrgK22AqJVlg2MMlD86/34d8f4071b6cc0d3b03c93f55e4d89b7/osi_model_session_layer_5.png)
  - This is the layer responsible for opening and closing communication between the two devices.
  - The time between when the communication is opened and closed is known as the session.
  - The session layer ensures that the session stays open long enough to transfer all the data being exchanged, and then promptly closes the session in order to avoid wasting resources.
  - The session layer also synchronizes data transfer with checkpoints. For example, if a 100 megabyte file is being transferred, the session layer could set a checkpoint every 5 megabytes. In the case of a disconnect or a crash after 52 megabytes have been transferred, the session could be resumed from the last checkpoint, meaning only 50 more megabytes of data need to be transferred. Without the checkpoints, the entire transfer would have to begin again from scratch.
- The transport layer (4)
  ![Transport Layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/3OlO75NcADGL3SmEADFDqd/723b8c7639c4e2e6b4febcbe7fd36e0e/osi_model_transport_layer_4.png)
  - Layer 4 is responsible for end-to-end communication between the two devices.
  - This includes taking data from the session layer and breaking it up into chunks called segments before sending it to layer 3. The transport layer on the receiving device is responsible for reassembling the segments into data the session layer can consume.
  - The transport layer is also responsible for flow control and error control. Flow control determines an optimal speed of transmission to ensure that a sender with a fast connection does not overwhelm a receiver with a slow connection.
  - The transport layer performs error control on the receiving end by ensuring that the data received is complete, and requesting a retransmission if it isn’t.
  - Transport layer protocols include the Transmission Control Protocol (TCP) and the User Datagram Protocol (UDP).
- The network layer (3)
  - ![Network Layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/3g2Hv0frHsql5SFauJL5EG/d8cede7b6a780e63413bd86de9eee7f9/osi_model_network_layer_3.png)
  - The network layer is responsible for facilitating data transfer between two different networks. If the two devices communicating are on the same network, then the network layer is unnecessary.
  - The network layer breaks up segments from the transport layer into smaller units, called packets, on the sender’s device, and reassembling these packets on the receiving device.
  - The network layer also finds the best physical path for the data to reach its destination; this is known as routing.
  - Network layer protocols include IP, the Internet Control Message Protocol (ICMP), the Internet Group Message Protocol (IGMP), and the IPsec suite.
- The data link layer (2)
  - ![Data link layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/3TLHavXiotb9ayyZFKECf3/9456d1c431cd71ceea7f4b407f076f11/data_link_layer_osi_model.png)
  - The data link layer is very similar to the network layer, except the data link layer facilitates data transfer between two devices on the same network.
  - The data link layer takes packets from the network layer and breaks them into smaller pieces called frames.
  - Like the network layer, the data link layer is also responsible for flow control and error control in intra-network communication (The transport layer only does flow control and error control for inter-network communications).
- The physical layer (1)
  - ![Physical Layer](https://cf-assets.www.cloudflare.com/slt3lc6tev37/1HQ1W5P4XAinIdM37DTu4U/900ccdceda346baf03ce8b9f977d2974/osi_model_physical_layer_1.png)
  - This layer includes the physical equipment involved in the data transfer, such as the cables and switches. 
  - This is also the layer where the data gets converted into a bit stream, which is a string of 1s and 0s.
  - The physical layer of both devices must also agree on a signal convention so that the 1s can be distinguished from the 0s on both devices.

### How data flows through the OSI Model
- In order for human-readable information to be transferred over a network from one device to another, the data must travel down the seven layers of the OSI Model on the sending device and then travel up the seven layers on the receiving end.
- For example: Mr. Cooper wants to send Ms. Palmer an email. Mr. Cooper composes his message in an email application on his laptop and then hits ‘send’. His email application will pass his email message over to the application layer, which will pick a protocol (SMTP) and pass the data along to the presentation layer. The presentation layer will then compress the data and then it will hit the session layer, which will initialize the communication session.
- The data will then hit the sender’s transportation layer where it will be segmented, then those segments will be broken up into packets at the network layer, which will be broken down even further into frames at the data link layer. The data link layer will then deliver those frames to the physical layer, which will convert the data into a bitstream of 1s and 0s and send it through a physical medium, such as a cable.
- Once Ms. Palmer’s computer receives the bit stream through a physical medium (such as her wifi), the data will flow through the same series of layers on her device, but in the opposite order. First the physical layer will convert the bitstream from 1s and 0s into frames that get passed to the data link layer. The data link layer will then reassemble the frames into packets for the network layer. The network layer will then make segments out of the packets for the transport layer, which will reassemble the segments into one piece of data.
- The data will then flow into the receiver's session layer, which will pass the data along to the presentation layer and then end the communication session. The presentation layer will then remove the compression and pass the raw data up to the application layer. The application layer will then feed the human-readable data along to Ms. Palmer’s email software, which will allow her to read Mr. Cooper’s email on her laptop screen.

## [OSI Model](https://www.imperva.com/learn/application-security/osi-model/) <span id="content-3"></span>

### What Is the OSI Model
- The Open Systems Interconnection (OSI) model describes seven layers that computer systems use to communicate over a network.
- It was the first standard model for network communications, adopted by all major computer and telecommunication companies in the early 1980s
- The modern Internet is not based on OSI, but on the simpler TCP/IP model.

### OSI Model Explained: The OSI 7 Layers
  ![Image](https://www.imperva.com/learn/wp-content/uploads/sites/13/2020/02/OSI-7-layers.jpg.webp)
- Application Layer
  - The application layer is used by end-user software such as web browsers and email clients. It provides protocols that allow software to send and receive information and present meaningful data to users. A few examples of application layer protocols are the Hypertext Transfer Protocol (HTTP), File Transfer Protocol (FTP), Post Office Protocol (POP), Simple Mail Transfer Protocol (SMTP), and Domain Name System (DNS).
- Presentation Layer
  - The presentation layer prepares data for the application layer. It defines how two devices should encode, encrypt, and compress data so it is received correctly on the other end. The presentation layer takes any data transmitted by the application layer and prepares it for transmission over the session layer.
- Session Layer
  - The session layer creates communication channels, called sessions, between devices. It is responsible for opening sessions, ensuring they remain open and functional while data is being transferred, and closing them when communication ends. The session layer can also set checkpoints during a data transfer—if the session is interrupted, devices can resume data transfer from the last checkpoint.
- Transport Layer
  - The transport layer takes data transferred in the session layer and breaks it into “segments” on the transmitting end. It is responsible for reassembling the segments on the receiving end, turning it back into data that can be used by the session layer. The transport layer carries out flow control, sending data at a rate that matches the connection speed of the receiving device, and error control, checking if data was received incorrectly and if not, requesting it again.
- Network Layer
  - The network layer has two main functions. One is breaking up segments into network packets, and reassembling the packets on the receiving end. The other is routing packets by discovering the best path across a physical network. The network layer uses network addresses (typically Internet Protocol addresses) to route packets to a destination node.
- Data Link Layer
  - The data link layer establishes and terminates a connection between two physically-connected nodes on a network. It breaks up packets into frames and sends them from source to destination. This layer is composed of two parts—Logical Link Control (LLC), which identifies network protocols, performs error checking and synchronizes frames, and Media Access Control (MAC) which uses MAC addresses to connect devices and define permissions to transmit and receive data.
- Physical Layer
  - The physical layer is responsible for the physical cable or wireless connection between network nodes. It defines the connector, the electrical cable or wireless technology connecting the devices, and is responsible for transmission of the raw data, which is simply a series of 0s and 1s, while taking care of bit rate control.


## [What Is Computer Networking?](https://aws.amazon.com/what-is/computer-networking/) <span id="content-4"></span>

### What is computer networking?
- Computer networking refers to interconnected computing devices that can exchange data and share resources with each other.
- These networked devices use a system of rules, called communications protocols, to transmit information over physical or wireless technologies.

### How does a computer network work?
- Nodes and links are the basic building blocks in computer networking.
- A network node may be data communication equipment (DCE) such as a modem, hub or, switch, or data terminal equipment (DTE) such as two or more computers and printers. 
- A link refers to the transmission media connecting two nodes. Links may be physical, like cable wires or optical fibers, or free space used by wireless networks.
- In a working computer network, nodes follow a set of rules or protocols that define how to send and receive electronic data via the links.
- The computer network architecture defines the design of these physical and logical components. It provides the specifications for the network’s physical components, functional organization, protocols, and procedures.

### What do computer networks do?
- Modern computer networks can:
  - Operate virtually
  - Integrate on a large scale
  - Respond quickly to changing conditions
  - Provide data security

### What are the types of computer network architecture?
- Client-server architecture
  - In this type of computer network, nodes may be servers or clients. Server nodes provide resources like memory, processing power, or data to client nodes. Server nodes may also manage client node behavior. Clients may communicate with each other, but they do not share resources. For example, some computer devices in enterprise networks store data and configuration settings. These devices are the servers in the network. Clients may access this data by making a request to the server machine.
- Peer-to-peer architecture
  - In Peer-to-Peer (P2P) architecture, connected computers have equal powers and privileges. There is no central server for coordination. Each device in the computer network can act as either client or server. Each peer may share some of its resources, like memory and processing power, with the entire computer network. For example, some companies use P2P architecture to host memory-consuming applications, such as 3-D graphic rendering, across multiple digital devices.

### What is network topology?
- Bus topology
  - Each node is linked to one other node only. Data transmission over the network connections occurs in one direction.
- Ring topology
  - Each node is linked to two other nodes, forming a ring. Data can flow bi-directionally. However,single node failure can bring down the entire network.
- Star topology
  - A central server node is linked to multiple client network devices. This topology performs better as data doesn’t have to go through each node. It is also more reliable.
- Mesh topology
  - Every node is connected to many other nodes. In a full mesh topology, every node is connected to every other node in the network.

### What are the types of enterprise computer networks?
- Local area network (LAN)
  - A LAN is an interconnected system limited in size and geography. It typically connects computers and devices within a single office or building. It is used by small companies or as a test network for small-scale prototyping.
- Wide area networks (WAN)
  - An enterprise network spanning buildings, cities, and even countries, is called a wide area network (WAN). While local area networks are used to transmit data at higher speeds within close proximity, WANs are set up for long-distance communication that is secure and dependable.
  - SD-WAN or software-defined WAN is virtual WAN architecture controlled by software technologies. An SD-WAN offers more flexible and dependable connectivity services that can be controlled at the application level without sacrificing security and quality of service.
- Service provider networks
  - Service provider networks allow customers to lease network capacity and functionality from the provider. Network service providers may consist of telecommunications companies, data carriers, wireless communications providers, Internet service providers, and cable television operators offering high-speed Internet access.
- Cloud networks
  - Conceptually, a cloud network can be seen as a WAN with its infrastructure delivered by a cloud-based service. Some or all of an organization’s network capabilities and resources are hosted in a public or private cloud platform and made available on demand. These network resources can include virtual routers, firewalls, bandwidth, and network management software,with other tools and functions available as required.
  - Businesses today use cloud networks to accelerate time-to-market, increase scale, and manage costs effectively. The cloud network model has become the standard approach for building and delivering applications for modern enterprises.


## [POSIX](https://www.techtarget.com/whatis/definition/POSIX-Portable-Operating-System-Interface#:~:text=POSIX%20(Portable%20Operating%20System%20Interface)%20is%20a%20set%20of%20standard,on%20the%20Unix%20operating%20system.) <span id="content-5"></span>

### What is POSIX (Portable Operating System Interface)?
- POSIX (Portable Operating System Interface) is a set of standard operating system interfaces based on the Unix operating system.
- POSIX emerged out of a need to make applications more portable across diverse systems
- In the early days of computing, programmers had to rewrite their applications for each computer model and OS.


## [What is the meaning of "POSIX"?](https://stackoverflow.com/questions/1780599/what-is-the-meaning-of-posix) <span id="content-6"></span>
- POSIX is a family of standards, specified by the IEEE, to clarify and make uniform the application programming interfaces (and ancillary issues, such as command line shell utilities) provided by Unix-y operating systems.
- When you write your programs to rely on POSIX standards, you can be pretty sure to be able to port them easily among a large family of Unix derivatives (including Linux, but not limited to it!); if and when you use some Linux API that's not standardized as part of Posix, you will have a harder time if and when you want to port that program or library to other Unix-y systems (e.g., MacOSX) in the future.


## [What is POSIX? Why Does it Matter to Linux/UNIX Users?](https://itsfoss.com/posix/) <span id="content-7"></span>

### What is POSIX?
- POSIX isn’t actually a thing. It describes a thing – much like a label. Imagine a box labeled: POSIX, and inside the box is a standard. A standard consists of sets of rules and instructions that POSIX is concerned with
- POSIX is shorthand for Portable Operating System Interface. It is an IEEE 1003.1 standard that defines the language interface between application programs (along with command line shells and utility interfaces) and the UNIX operating system.
- Compliance to the standard ensures compatibility when UNIX programs are moved from one UNIX platform to another.
- A standard must be spelled out and followed by rules on how to achieve the goal of interoperability between operating systems.

### Why POSIX?
- It’s important to note that POSIX wasn’t formed to control how the operating systems were built – any company was free to design their UNIX variant any way they pleased. POSIX was only concerned with how an application interfaces with the operating system.
- In programmer-speak, an interface is the method one program’s code can communicate with another program. The interface expects Program A to provide a specific type of information to Program B. Likewise, Program A expects Program B to answer back with a specific type of data.
- Without going into a lot of programmer-speak, I’ll just say that the cat command makes a call to the operating system to fetch the file so cat can read it. cat reads it and then displays the file’s contents on the screen. There is a lot of interplay between the application (cat) and the operating system. How this interplay works is what POSIX was interested in.
-  If the interplay could be the same across the different UNIX variants, portability – regardless of operating system, manufacturer, and hardware – is regained.

### Compliance is Voluntary
- When code is compliant, it’s easier to move to another system; very little code rewrite, if any, would be necessary.
- When code can work on different systems, the use of it expands. People using other systems can benefit from the use of the program.
- When operating systems and programs conform to the POSIX standard, they gain the benefit of interoperability. They will be able to move from one system to another with the reasonable expectation that the machines will work much like another one does. Their data will still be accessible and they will still be able to make changes to it.

### Conclusion
- The POSIX standard allows developers to create applications, tools, and platforms on many operating systems using much of the same code.
- Basically, POSIX is geared toward operating system designers and software developers, but as users of a system


## [What is a Network Socket?](https://www.tutorialspoint.com/what-is-a-network-socket-computer-networks) <span id="content-8"></span>

### Introduction
- A network socket is a software component within a computer network node that acts as an endpoint for delivering and receiving data.
- In this context, a socket's address, which is the triad of the transport protocol, IP address, and port number, is used to externally identify it to other hosts.

### Types of Sockets
- Datagram Socket
  - A datagram socket is a type of network socket in which packets are sent and received without the use of a link. It resembles a mailbox. Letters (data) are gathered and delivered (transmitted) to a mailbox (receiving socket). It is a connection-less socket.
- Stream Socket
  - A stream socket is a type of network socket in a computer operating system that provides a connection-oriented, sequenced, and unique flow of data without record boundaries, as well as well-defined mechanisms for creating and destroying connections and detecting errors. It is comparable to a telephone. Between the phones (two ends), a link is created. 
- The following diagram shows the complete Client and Server interaction
  ![Diagram -- Client and Server Interaction](https://www.tutorialspoint.com/assets/questions/media/55610/types_of_sockets.jpg)


## [What is a Socket?](https://docs.oracle.com/javase/tutorial/networking/sockets/definition.html#:~:text=Definition%3A,address%20and%20a%20port%20number.) <span id="content-9"></span>
- The server just waits, listening to the socket for a client to make a connection request.
- On the client-side: The client knows the hostname of the machine on which the server is running and the port number on which the server is listening. To make a connection request, the client tries to rendezvous with the server on the server's machine and port.
- The client also needs to identify itself to the server so it binds to a local port number that it will use during this connection. This is usually assigned by the system.
- If everything goes well, the server accepts the connection. Upon acceptance, the server gets a new socket bound to the same local port and also has its remote endpoint set to the address and port of the client.
- It needs a new socket so that it can continue to listen to the original socket for connection requests while tending to the needs of the connected client.
- On the client side, if the connection is accepted, a socket is successfully created and the client can use the socket to communicate with the server.
- Definition:
  - A socket is one endpoint of a two-way communication link between two programs running on the network. A socket is bound to a port number so that the TCP l
- An endpoint is a combination of an IP address and a port number. Every TCP connection can be uniquely identified by its two endpoints.


## [Socket Programming in Computer Network](https://www.scaler.com/topics/computer-network/socket-programming/) <span id="content-10"></span>

### Overview
- Sockets in computer networks are used for allowing the transmission of information between two processes of the same machines or different machines in the network.
- The socket is the combination of IP address and software port number used for communication between multiple processes. 
- Socket helps to recognize the address of the application to which data is to be sent using the IP address and port number.

### What is Socket Programming in Computer Networks?
- Sockets are the end of two-way communication between two programs that are running on the networks.
- Sockets are mostly used in client-server architecture for communication between multiple applications.
- Socket address:
  ![Socket Address](https://scaler.com/topics/images/socket-programming-in-computer-networks.webp)
- TCP Socket Connection:
  ![TCP Socket Connection](https://scaler.com/topics/images/socket-programming-in-tcp.webp)


## [A quick introduction to processes in Computer Science](https://medium.com/@imdadahad/a-quick-introduction-to-processes-in-computer-science-271f01c780da) <span id="content-11"></span>

### The basic idea
- A process is simply a program in execution.

### Monitoring processes in the system
- At any given time there may be a couple hundred or less processes running. Sometimes it useful to see detailed information about them, especially if your computer is running slow: as certain processes may be hogging the computer’s memory or the CPU.

### How does the operating system manage processes
- Process:
  ![Process](https://miro.medium.com/v2/resize:fit:1304/format:webp/1*LdafsmbCdfR882Vtd1Kp0g.jpeg)
- To do this effectively, the operating system maintains what’s known as a Process Control Block for each process. It contains useful information such as the current process state, the next instruction to perform and currently allocated devices to the process.


## [Operating System - Processes](https://www.tutorialspoint.com/operating_system/os_processes.htm) <span id="content-12"></span>

### Process
- A process is basically a program in execution. The execution of a process must progress in a sequential fashion.
- A process is defined as an entity which represents the basic unit of work to be implemented in the system.

### Program
- A program is a piece of code which may be a single line or millions of lines.
- A computer program is a collection of instructions that performs a specific task when executed by a computer.
- When we compare a program with a process, we can conclude that a process is a dynamic instance of a computer program.

### Process Life Cycle
- Diagram:
  ![Process Life Cycle - Diagram](https://www.tutorialspoint.com/operating_system/images/process_state.jpg)

### Process Control Block (PCB)
- A Process Control Block is a data structure maintained by the Operating System for every process.
- The PCB is identified by an integer process ID (PID). A PCB keeps all the information needed to keep track of a process as listed below in the table


## [I/O Hardware](https://www.tutorialspoint.com/operating_system/os_io_hardware.htm) <span id="content-13"></span>
- Block devices − A block device is one with which the driver communicates by sending entire blocks of data. For example, Hard disks, USB cameras, Disk-On-Key etc.
- Character devices − A character device is one with which the driver communicates by sending and receiving single characters (bytes, octets). For example, serial ports, parallel ports, sounds cards etc
- Synchronous I/O − In this scheme CPU execution waits while I/O proceeds
- Asynchronous I/O − I/O proceeds concurrently with CPU execution


## [Virtualization](https://www.techtarget.com/searchitoperations/definition/virtualization) <span id="content-14"></span>

### What is virtualization?
- Virtualization is the creation of a virtual -- rather than actual -- version of something, such as an operating system (OS), a server, a storage device or network resources.
- Virtualization uses software that simulates hardware functionality to create a virtual system.
- OS virtualization is the use of software to allow a piece of hardware to run multiple operating system images at the same time.

### How virtualization works
- A key use of virtualization technology is server virtualization, which uses a software layer -- called a hypervisor
- Hypervisors take the physical resources and separate them so they can be utilized by the virtual environment.
- They can sit on top of an OS or they can be directly installed onto the hardware. The latter is how most enterprises virtualize their systems.
- The Xen hypervisor is an open source software program that is responsible for managing the low-level interactions that occur between virtual machines (VMs) and the physical hardware.
- In other words, the Xen hypervisor enables the simultaneous creation, execution and management of various virtual machines in one physical environment.
- With the help of the hypervisor, the guest OS, normally interacting with true hardware, is now doing so with a software emulation of that hardware; often, the guest OS has no idea it's on virtualized hardware.
- Traditional and virtual architecture:
  ![Traditional and virtual architecture](https://cdn.ttgtmedia.com/rms/onlineImages/server_virtualization-traditional_virtual_architecture.jpg)
- The virtualization process follows the steps listed below:
  - Hypervisors detach the physical resources from their physical environments.
  - Resources are taken and divided, as needed, from the physical environment to the various virtual environments.
  - System users work with and perform computations within the virtual environment.
  - Once the virtual environment is running, a user or program can send an instruction that requires extra resources form the physical environment. In response, the hypervisor relays the message to the physical system and stores the changes. This process will happen at an almost native speed.
- The VM acts like a single data file that can be transferred from one computer to another and opened in both; it is expected to perform the same way on every computer.

### Types of virtualization
- A partition is the logical division of a hard disk drive to create, in effect, two separate hard drives.
- There are six areas of IT where virtualization is making headway:
  - Network virtualization is a method of combining the available resources in a network by splitting up the available bandwidth into channels, each of which is independent from the others and can be assigned -- or reassigned -- to a particular server or device in real time. The idea is that virtualization disguises the true complexity of the network by separating it into manageable parts, much like your partitioned hard drive makes it easier to manage your files.
  - Storage virtualization is the pooling of physical storage from multiple network storage devices into what appears to be a single storage device that is managed from a central console. Storage virtualization is commonly used in storage area networks.
  - Server virtualization is the masking of server resources -- including the number and identity of individual physical servers, processors and operating systems -- from server users. The intention is to spare the user from having to understand and manage complicated details of server resources while increasing resource sharing and utilization and maintaining the capacity to expand later.
    - The layer of software that enables this abstraction is often referred to as the hypervisor.
    - The most common hypervisor -- Type 1 -- is designed to sit directly on bare metal and provide the ability to virtualize the hardware platform for use by the virtual machines.
    -  KVM virtualization is a Linux kernel-based virtualization hypervisor that provides Type 1 virtualization benefits like other hypervisors. KVM is licensed under open source. A Type 2 hypervisor requires a host operating system and is more often used for testing and labs.
  - Data virtualization is abstracting the traditional technical details of data and data management, such as location, performance or format, in favor of broader access and more resiliency tied to business needs.
  - Desktop virtualization is virtualizing a workstation load rather than a server. This allows the user to access the desktop remotely, typically using a thin client at the desk. Since the workstation is essentially running in a data center server, access to it can be both more secure and portable. The operating system license does still need to be accounted for as well as the infrastructure.
  - Application virtualization is abstracting the application layer away from the operating system. This way, the application can run in an encapsulated form without being depended upon on by the operating system underneath. This can allow a Windows application to run on Linux and vice versa, in addition to adding a level of isolation.

## [What is virtualization?](https://opensource.com/resources/virtualization) <span id="content-15"></span>
- A hypervisor is a program for creating and running virtual machines.
- Type one, or "bare metal" hypervisors that run guest virtual machines directly on a system's hardware, essentially behaving as an operating system
- Type two, or "hosted" hypervisors behave more like traditional applications that can be started and stopped like a normal program.
- A virtual machine is the emulated equivalent of a computer system that runs on top of another system.
- A container is actually just an isolated process that shared the same Linux kernel as the host operating system, as well as the libraries and other files needed for the execution of the program running inside of the container
- Typically, containers are designed to run a single program, as opposed to emulating a full multi-purpose server.


## [Containers vs. virtual machines](https://www.atlassian.com/microservices/cloud-computing/containers-vs-vms) <span id="content-16"></span>
- Virtualization is the process in which a system singular resource like RAM, CPU, Disk, or Networking can be ‘virtualized’ and represented as multiple resources.
- The key differentiator between containers and virtual machines is that virtual machines virtualize an entire machine down to the hardware layers and containers only virtualize software layers above the operating system level.
- Image:
  ![Image](https://wac-cdn.atlassian.com/dam/jcr:92adde69-f728-4cfc-8bab-ba391c25ae58/SWTM-2060_Diagram_Containers_VirtualMachines_v03.png?cdnVersion=1040)

### What is a container?
- Containers are lightweight software packages that contain all the dependencies required to execute the contained software application. These dependencies include things like system libraries, external third-party code packages, and other operating system level applications.
- Pros
  - Iteration speed
    - Because containers are lightweight and only include high level software, they are very fast to modify and iterate on.
  - Robust ecosystem
    - Most container runtime systems offer a hosted public repository of pre-made containers. These container repositories contain many popular software applications like databases or messaging systems and can be instantly downloaded and executed, saving time for development teams
- Cons
  - Shared host exploits
    - Containers all share the same underlying hardware system below the operating system layer, it is possible that an exploit in one container could break out of the container and affect the shared hardware. Most popular container runtimes have public repositories of pre-built containers. There is a security risk in using one of these public images as they may contain exploits or may be vulnerable to being hijacked by nefarious actors.
 
### What is a virtual machine?
- Virtual machines are heavy software packages that provide complete emulation of low level hardware devices like CPU, Disk and Networking devices
- Pros
  - Full isolation security
    - Virtual machines run in isolation as  a fully standalone system.
  - Interactive development
    - Virtual machines are more dynamic and can be interactively developed. Once the basic hardware definition is specified for a virtual machine the virtual machine can then be treated as a bare bones computer.
- Cons
  - Iteration speed
  - Storage size cost


## [File System](https://www.tutorialspoint.com/operating_system/os_file_system.htm) <span id="content-17"></span>
- A file is a named collection of related information that is recorded on secondary storage such as magnetic disks, magnetic tapes and optical disks. In general, a file is a sequence of bits, bytes, lines or records whose meaning is defined by the files creator and user.
- File type refers to the ability of the operating system to distinguish different types of file such as text files source files and binary files etc.
- A sequential access is that in which the records are accessed in some sequence, i.e., the information in the file is processed in order, one record after the other. This access method is the most primitive one. Example: Compilers usually access files in this fashion.
- Direct/Random access
  - Random access file organization provides, accessing the records directly.
  - Each record has its own address on the file with by the help of which it can be directly accessed for reading or writing.
  - The records need not be in any sequence within the file and they need not be in adjacent locations on the storage medium.
- Indexed sequential access
  - This mechanism is built up on base of sequential access.
  - An index is created for each file which contains pointers to various blocks.
  - Index is searched sequentially and its pointer is used to access the file directly.
- Space Allocation
  - Contiguous Allocation
    - Each file occupies a contiguous address space on disk.
    - Assigned disk address is in linear order.
    - Easy to implement.
    - External fragmentation is a major issue with this type of allocation technique.
  - Linked Allocation
    - Each file carries a list of links to disk blocks.
    - Directory contains link / pointer to first block of a file.
    - No external fragmentation
    - Effectively used in sequential access file.
    - Inefficient in case of direct access file.
  - Indexed Allocation
    - Provides solutions to problems of contiguous and linked allocation.
    - A index block is created having all pointers to files.
    - Each file has its own index block which stores the addresses of disk space occupied by the file.
    - Directory contains the addresses of index blocks of files.


## [Demystifying memory management in modern programming languages](https://dev.to/deepu105/demystifying-memory-management-in-modern-programming-languages-ddd) <span id="content-18"></span>
- Memory management is the process of controlling and coordinating the way a software application access computer memory.
- When a software runs on a target Operating system on a computer it needs access to the computers RAM(Random-access memory) to:
  - load its own bytecode that needs to be executed
  - store the data values and data structures used by the program that is executed
  - load any run-time systems that are required for the program to execute
- The stack is used for static memory allocation and as the name suggests it is a last in first out(LIFO) stack
  - Due to this nature, the process of storing and retrieving data from the stack is very fast as there is no lookup required, you just store and retrieve data from the topmost block on it.
  - But this means any data that is stored on the stack has to be finite and static(The size of the data is known at compile-time).
  - This is where the execution data of the functions are stored as stack frames(So, this is the actual execution stack). Each frame is a block of space where the data required for that function is stored. For example, every time a function declares a new variable, it is "pushed" onto the topmost block in the stack. Then every time a function exits, the topmost block is cleared, thus all of the variables pushed onto the stack by that function, are cleared. These can be determined at compile time due to the static nature of the data stored here.
- Multi-threaded applications can have a stack per thread.
- Memory management of the stack is simple and straightforward and is done by the OS.
- Typical data that are stored on stack are local variables(value types or primitives, primitive constants), pointers and function frames.
- This is where you would encounter stack overflow errors as the size of the stack is limited compared to the Heap.
- There is a limit on the size of value that can be stored on the Stack for most languages.
- Heap is used for dynamic memory allocation and unlike stack, the program needs to look up the data in heap using pointers
  - It is slower than stack as the process of looking up data is more involved but it can store more data than the stack.
  - This means data with dynamic size can be stored here.
  - Heap is shared among threads of an application.
  - Due to its dynamic nature heap is trickier to manage and this is where most of the memory management issues arise from and this is where the automatic memory management solutions from the language kick in.
  - Typical data that are stored on the heap are global variables, reference types like objects, strings, maps, and other complex data structures.
  - This is where you would encounter out of memory errors if your application tries to use more memory than the allocated heap(Though there are many other factors at play here like GC, compacting).
  - Generally, there is no limit on the size of the value that can be stored on the heap. Of course, there is the upper limit of how much memory is allocated to the application.
- Unlike Hard disk drives, RAM is not infinite. If a program keeps on consuming memory without freeing it, ultimately it will run out of memory and crash itself or even worse crash the operating system.
- Manual memory management
  - For example, C and C++.
- Garbage collection(GC)
  - JVM(Java/Scala/Groovy/Kotlin), JavaScript, C#, Golang, OCaml, and Ruby are some of the languages that use Garbage collection for memory management by default.
  - Mark & Sweep GC: Also known as Tracing GC. Its generally a two-phase algorithm that first marks objects that are still being referenced as "alive" and in the next phase frees the memory of objects that are not alive. JVM, C#, Ruby, JavaScript, and Golang employ this approach for example. In JVM there are different GC algorithms to choose from while JavaScript engines like V8 use a Mark & Sweep GC along with Reference counting GC to complement it. This kind of GC is also available for C & C++ as an external library.
    ![Process of Garbage Collection](https://res.cloudinary.com/practicaldev/image/fetch/s--JxvXuUl1--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_66%2Cw_880/https://i.imgur.com/AZaR0LP.gif)
  - Reference counting GC: In this approach, every object gets a reference count which is incremented or decremented as references to it change and garbage collection is done when the count becomes zero. It's not very preferred as it cannot handle cyclic references. PHP, Perl, and Python, for example, uses this type of GC with workarounds to overcome cyclic references. This type of GC can be enabled for C++ as well.
- Resource Acquisition is Initialization (RAII)
  - In this type of memory management, an object's memory allocation is tied to its lifetime, which is from construction until destruction. It was introduced in C++ and is also used by Ada and Rust.
  - Automatic Reference Counting(ARC)
    - It's similar to Reference counting GC but instead of running a runtime process at a specific interval the retain and release instructions are inserted to the compiled code at compile-time and when an object reference becomes zero its cleared automatically as part of execution without any program pause. It also cannot handle cyclic references and relies on the developer to handle that by using certain keywords. Its a feature of the Clang compiler and provides ARC for Objective C & Swift.
  - Ownership
    - It combines RAII with an ownership model, any value must have a variable as its owner(and only one owner at a time) when the owner goes out of scope the value will be dropped freeing the memory regardless of it being in stack or heap memory. It is kind of like Compile-time reference counting. It is used by Rust, in my research I couldn't find any other language using this exact mechanism.


## [Operating System: Threads and Concurrency](https://medium.com/@akhandmishra/operating-system-threads-and-concurrency-aec2036b90f8) <span id="content-19"></span>

### Introduction
- Definition: A Thread also called lightweight process, is basic CPU utilization; it compromises a thread ID, a program counter, a register set, and a stack. A thread is an entity within a process that can be scheduled for execution.
- If we want a process to be able to execute on multiple CPUs at a time, to take advantage of the multi-core systems, the process must have several execution-context called threads.
- A thread is an active entity which executes a part of a process.
-  Multiple threads execute simultaneously with each other which results in the execution of a single whole process. 
- The threads of a process are the part of the same virtual address space (i.e. they share all virtual to physical mapping), they share all the data, code and file.
- However, they will be executing the different instruction, they will be accessing the different portion of the address space or different in other ways.
- A different thread has a different stack, a different stack pointer register, a different Program Counter, and other registers.
- PCB: Single-threaded and multi-threaded process
  ![PCB: Single-threaded and multi-threaded process](https://miro.medium.com/v2/resize:fit:1280/format:webp/0*cgAbgEkofShV4B5z.png)

### Benefits of Multi-threading
- Parallelization: In multi-processor architecture, different threads can execute different instructions at a time, which result in parallelization which speeds up the execution of the process.
- Specialization: By specializing different threads to perform the different task, we can manage threads, for example, we can give higher priority to those threads which are executing the more important task. Also in multi-processor architecture specialization leads to the hotter cache which improves performance.
- Efficient: Multi-threaded programs are more efficient as compared to multi-process programs in terms of resource requirement as threads share address space while process does not, so multi-process programs will require separate memory space allocation. Also, Multi-threaded programs incur lower overhead as inter-thread communication is less expensive.
- Hide Latency: As context switching among the process is a costly operation, as all the threads of a process share the same virtual to physical address mapping and other resources, context switch among the thread is a cheap operation and require less time. When the number of thread is greater than the number of CPU and a thread is in idle state (spending the time to wait for the result of some interrupt) and its idle time is greater than two times the time required for switching the context to another thread, it will switch the switch context to another thread to hide idling time.

### Synchronisation Mechanisms:
- To deal with concurrency issues a mechanism is needed to execute threads in an exclusive manner to ensure threads access data and other resources one at a time, for this, we use a mutex which is nothing but mutual exclusion object which allows multiple threads to share resources like file access or memory access, but not simultaneously.
- A waiting mechanism is also needed for threads which are waiting for other threads to complete, specifying what are they waiting for so that they are not required to keeps on checking whether they are allowed to execute the operation or not, they will be notified whenever they are allowed. This type of inter-thread coordination is generally handled by condition variable.

### Thread and Thread Creation
- A thread data structure contains information about thread identity, register values like program counter, stack pointer, etc, stack and other attributes which help thread management system to manage threads like scheduling threads, debugging threads, etc.
- A mutex is like a lock which is used whenever the thread is accessing the data or resources that are shared among different threads.
- When a thread locks a mutex it has exclusive access to the shared resources. Other threads attempting to lock the mutex (as other threads may also need to perform operations that require exclusive access to shared resources) are not going to succeed and have to wait till the thread which has locked the mutex (i.e. the owner thread) completes its task.
- The data structure of mutex at least contains information about its lock status (whether the mutex is locked or not), list of all the blocked threads which are waiting for the mutex to be free, i.e. the owner thread to complete its work.
- The portion of the code performed by the thread under the locked state of the mutex is called critical section.
- The critical section contains the code which performs the operations which require only one thread at a time to perform. 
- When the owner thread exits the critical section it releases the mutex and other waiting thread can lock it for their exclusive access to shared resources.



**[⬆ back to top](#list-of-contents)**

<br />

---

## References:
- https://roadmap.sh/devops
- https://www.cloudflare.com/learning/ddos/glossary/open-systems-interconnection-model-osi/
- https://www.imperva.com/learn/application-security/osi-model/
- https://aws.amazon.com/what-is/computer-networking/
- https://www.techtarget.com/whatis/definition/POSIX-Portable-Operating-System-Interface#:~:text=POSIX%20(Portable%20Operating%20System%20Interface)%20is%20a%20set%20of%20standard,on%20the%20Unix%20operating%20system.
- https://stackoverflow.com/questions/1780599/what-is-the-meaning-of-posix
- https://www.baeldung.com/linux/posix
- https://itsfoss.com/posix/
