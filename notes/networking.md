# Networking

<br />

## List of Contents:
### 1. [Roadmap.sh - Networking](#content-1)
### 2. [What is the OSI Model?](#content-2)
### 3. [OSI Model](#content-3)
### 4. [What Is Computer Networking?](#content-4)

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


**[⬆ back to top](#list-of-contents)**

<br />

---

## References:
- https://roadmap.sh/devops
- https://www.cloudflare.com/learning/ddos/glossary/open-systems-interconnection-model-osi/
- https://www.imperva.com/learn/application-security/osi-model/
- https://aws.amazon.com/what-is/computer-networking/
