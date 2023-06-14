# Web Servers

<br />

## List of Contents:
### 1. [What is a reverse proxy? | Proxy servers explained](#content-1)
### 2. [What is a Proxy Server? How does it work?](#content-2)


<br />

---

## Contents:

## [What is a reverse proxy? | Proxy servers explained](https://www.cloudflare.com/en-gb/learning/cdn/glossary/reverse-proxy/) <span id="content-1"></span>

### What is a reverse proxy?
- A reverse proxy is a server that sits in front of web servers and forwards client (e.g. web browser) requests to those web servers.
- Reverse proxies are typically implemented to help increase security, performance, and reliability.
- In order to better understand how a reverse proxy works and the benefits it can provide, let’s first define what a proxy server is.

### How is a reverse proxy different?
- A reverse proxy is a server that sits in front of one or more web servers, intercepting requests from clients.
- With a reverse proxy, when clients send requests to the origin server of a website, those requests are intercepted at the network edge by the reverse proxy server.
- The reverse proxy server will then send requests to and receive responses from the origin server.
- Once again, let’s illustrate by naming the computers involved:
  - D: Any number of users’ home computers
  - E: This is a reverse proxy server
  - F: One or more origin servers
  - Flow:
    ![Reverse Proxy Flow](https://cf-assets.www.cloudflare.com/slt3lc6tev37/3msJRtqxDysQslvrKvEf8x/f7f54c9a2cad3e4586f58e8e0e305389/reverse_proxy_flow.png)
- Typically all requests from D would go directly to F, and F would send responses directly to D. With a reverse proxy, all requests from D will go directly to E, and E will send its requests to and receive responses from F. E will then pass along the appropriate responses to D.
- Below we outline some of the benefits of a reverse proxy:
  - Load balancing - A popular website that gets millions of users every day may not be able to handle all of its incoming site traffic with a single origin server. Instead, the site can be distributed among a pool of different servers, all handling requests for the same site. In this case, a reverse proxy can provide a load balancing solution which will distribute the incoming traffic evenly among the different servers to prevent any single server from becoming overloaded. In the event that a server fails completely, other servers can step up to handle the traffic.
  - Protection from attacks - With a reverse proxy in place, a web site or service never needs to reveal the IP address of their origin server(s). This makes it much harder for attackers to leverage a targeted attack against them, such as a DDoS attack. Instead the attackers will only be able to target the reverse proxy, such as Cloudflare’s CDN, which will have tighter security and more resources to fend off a cyber attack.
  - Global server load balancing (GSLB) - In this form of load balancing, a website can be distributed on several servers around the globe and the reverse proxy will send clients to the server that’s geographically closest to them. This decreases the distances that requests and responses need to travel, minimizing load times.
  - Caching - A reverse proxy can also cache content, resulting in faster performance. For example, if a user in Paris visits a reverse-proxied website with web servers in Los Angeles, the user might actually connect to a local reverse proxy server in Paris, which will then have to communicate with an origin server in L.A. The proxy server can then cache (or temporarily save) the response data. Subsequent Parisian users who browse the site will then get the locally cached version from the Parisian reverse proxy server, resulting in much faster performance.
  - SSL encryption - Encrypting and decrypting SSL (or TLS) communications for each client can be computationally expensive for an origin server. A reverse proxy can be configured to decrypt all incoming requests and encrypt all outgoing responses, freeing up valuable resources on the origin server.
  - 

[⬆ back to top](#list-of-contents)**

<br />

---

## [What is a Proxy Server? How does it work?](https://www.fortinet.com/resources/cyberglossary/proxy-server) <span id="content-2"></span>

### Proxy Server Definition
- A proxy server is a system or router that provides a gateway between users and the internet. Therefore, it helps prevent cyber attackers from entering a private network. It is a server, referred to as an “intermediary” because it goes between end-users and the web pages they visit online.
- When a computer connects to the internet, it uses an IP address. This is similar to your home’s street address, telling incoming data where to go and marking outgoing data with a return address for other devices to authenticate. A proxy server is essentially a computer on the internet that has an IP address of its own.

### Proxy Servers and Network Security
- Proxies provide a valuable layer of security for your computer. They can be set up as web filters or firewalls, protecting your computer from internet threats like malware.
- This extra security is also valuable when coupled with a secure web gateway or other email security products. This way, you can filter traffic according to its level of safety or how much traffic your network—or individual computers—can handle.
- Big picture:
  ![A Proxy Server in Action](https://www.fortinet.com/content/fortinet-com/en_us/resources/cyberglossary/proxy-server/_jcr_content/par/c05_container_copy_c/par/c28_image_copy_copy_.img.jpg/1625683502431.jpg)

### How a Proxy Works
- Because a proxy server has its own IP address, it acts as a go-between for a computer and the internet. Your computer knows this address, and when you send a request on the internet, it is routed to the proxy, which then gets the response from the web server and forwards the data from the page to your computer’s browser, like Chrome, Safari, Firefox, or Microsoft Edge

### Benefits of a Proxy Server
- Enhanced security: Can act like a firewall between your systems and the internet. Without them, hackers have easy access to your IP address, which they can use to infiltrate your computer or network.
- Private browsing, watching, listening, and shopping: Use different proxies to help you avoid getting inundated with unwanted ads or the collection of IP-specific data. With a proxy, site browsing is well-protected and impossible to track.
- Access to location-specific content: You can designate a proxy server with an address associated with another country. You can, in effect, make it look like you are in that country and gain full access to all the content computers in that country are allowed to interact with. For example, the technology can allow you to open location-restricted websites by using local IP addresses of the location you want to appear to be in. 
- Prevent employees from browsing inappropriate or distracting sites: You can use it to block access to websites that run contrary to your organization’s principles. Also, you can block sites that typically end up distracting employees from important tasks. Some organizations block social media sites like Facebook and others to remove time-wasting temptations.

### Types of Proxy Servers
- Forward Proxy 
  - A forward proxy sits in front of clients and is used to get data to groups of users within an internal network. When a request is sent, the proxy server examines it to decide whether it should proceed with making a connection.
  - A forward proxy is best suited for internal networks that need a single point of entry. It provides IP address security for those in the network and allows for straightforward administrative control. However, a forward proxy may limit an organization’s ability to cater to the needs of individual end-users.
- Transparent Proxy
  - A transparent proxy can give users an experience identical to what they would have if they were using their home computer. In that way, it is “transparent.” They can also be “forced” on users, meaning they are connected without knowing it.
  - Transparent proxies are well-suited for companies that want to make use of a proxy without making employees aware they are using one. It carries the advantage of providing a seamless user experience. On the other hand, transparent proxies are more susceptible to certain security threats, such as SYN-flood denial-of-service attacks.
- Anonymous Proxy
  - An anonymous proxy focuses on making internet activity untraceable. It works by accessing the internet on behalf of the user while hiding their identity and computer information.
  - A anonymous proxy is best suited for users who want to have full anonymity while accessing the internet. While anonymous proxies provide some of the best identity protection possible, they are not without drawbacks. Many view the use of anonymous proxies as underhanded, and users sometimes face pushback or discrimination as a result.
- High Anonymity Proxy
  - A high anonymity proxy is an anonymous proxy that takes anonymity one step further. It works by erasing your information before the proxy attempts to connect to the target site.
  - The server is best suited for users for whom anonymity is an absolute necessity, such as employees who do not want their activity traced back to the organization. On the downside, some of them, particularly the free ones, are decoys set up to trap users in order to access their personal information or data.
- Distorting Proxy
  - A distorting proxy identifies itself as a proxy to a website but hides its own identity. It does this by changing its IP address to an incorrect one.
  - Distorting proxies are a good choice for people who want to hide their location while accessing the internet. This type of proxy can make it look like you are browsing from a specific country and give you the advantage of hiding not just your identity but that of the proxy, too. This means even if you are associated with the proxy, your identity is still secure. However, some websites automatically block distorting proxies, which could keep an end-user from accessing sites they need.
- Data Center Proxy
  - Data center proxies are not affiliated with an internet service provider (ISP) but are provided by another corporation through a data center. The proxy server exists in a physical data center, and the user’s requests are routed through that server.
  - Data center proxies are a good choice for people who need quick response times and an inexpensive solution. They are therefore a good choice for people who need to gather intelligence on a person or organization very quickly. They carry the benefit of giving users the power to swiftly and inexpensively harvest data. On the other hand, they do not offer the highest level of anonymity, which may put users’ information or identity at risk.
- Residential Proxy
  - A residential proxy gives you an IP address that belongs to a specific, physical device. All requests are then channeled through that device.
  - Residential proxies are well-suited for users who need to verify the ads that go on their website, so you can block cookies, suspicious or unwanted ads from competitors or bad actors. Residential proxies are more trustworthy than other proxy options. However, they often cost more money to use, so users should carefully analyze whether the benefits are worth the extra investment.
- Public Proxy
  - A public proxy is accessible by anyone free of charge. It works by giving users access to its IP address, hiding their identity as they visit sites. 
  - Public proxies are best suited for users for whom cost is a major concern and security and speed are not. Although they are free and easily accessible, they are often slow because they get bogged down with free users. When you use a public proxy, you also run an increased risk of having your information accessed by others on the internet.
- Shared Proxy
  - Shared proxies are used by more than one user at once. They give you access to an IP address that may be shared by other people, and then you can surf the internet while appearing to browse from a location of your choice.
  - Shared proxies are a solid option for people who do not have a lot of money to spend and do not necessarily need a fast connection. The main advantage of a shared proxy is its low cost. Because they are shared by others, you may get blamed for someone else’s bad decisions, which could get you banned from a site.
- SSL Proxy
  - A secure sockets layer (SSL) proxy provides decryption between the client and the server. As the data is encrypted in both directions, the proxy hides its existence from both the client and the server.
  - These proxies are best suited for organizations that need enhanced protection against threats that the SSL protocol reveals and stops. Because Google prefers servers that use SSL, an SSL proxy, when used in connection with a website, may help its search engine ranking. On the downside, content encrypted on an SSL proxy cannot be cached, so when visiting websites multiple times, you may experience slower performance than you would otherwise.
- Rotating Proxy
  - A rotating proxy assigns a different IP address to each user that connects to it. As users connect, they are given an address that is unique from the device that connected before it.
  - Rotating proxies are ideal for users who need to do a lot of high-volume, continuous web scraping. They allow you to return to the same website again and again anonymously. However, you have to be careful when choosing rotating proxy services. Some of them contain public or shared proxies that could expose your data.


[⬆ back to top](#list-of-contents)**

<br />

---

## References:
- https://www.cloudflare.com/en-gb/learning/cdn/glossary/reverse-proxy/
