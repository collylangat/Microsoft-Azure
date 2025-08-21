# Introduction to cloud computing
# What is Cloud Computing

Cloud computing is the delivery of computing services over the internet.
Computing services include common IT infrastructure such as virtual machines, storage, databases, and networking
Because cloud computing uses the internet to deliver these services, it doesn’t have to be constrained by physical infrastructure the same way that a traditional datacenter is. That means if you need to increase your IT infrastructure rapidly, you don’t have to wait to build a new datacenter—you can use the cloud to rapidly expand your IT footprint.

# Shared responsiblity model
With the shared responsibility model, these responsibilities get shared between the cloud provider and the consumer. Physical security, power, cooling, and network connectivity are the responsibility of the cloud provider. The consumer isn’t collocated with the datacenter, so it wouldn’t make sense for the consumer to have any of those responsibilities.

At the same time, the consumer is responsible for the data and information stored in the cloud. (You wouldn’t want the cloud provider to be able to read your information.) The consumer is also responsible for access security, meaning you only give access to those who need it.

Then, for some things, the responsibility depends on the situation. If you’re using a cloud SQL database, the cloud provider would be responsible for maintaining the actual database. However, you’re still responsible for the data that gets ingested into the database. If you deployed a virtual machine and installed an SQL database on it, you’d be responsible for database patches and updates, as well as maintaining the data and information stored in the database.

With an on-premises datacenter, you’re responsible for everything. With cloud computing, those responsibilities shift. The shared responsibility model is heavily tied into the cloud service types (covered later in this learning path): infrastructure as a service (IaaS), platform as a service (PaaS), and software as a service (SaaS). IaaS places the most responsibility on the consumer, with the cloud provider being responsible for the basics of physical security, power, and connectivity. On the other end of the spectrum, SaaS places most of the responsibility with the cloud provider. PaaS, being a middle ground between IaaS and SaaS, rests somewhere in the middle and evenly distributes responsibility between the cloud provider and the consumer.

When using a cloud provider, you’ll always be responsible for:

- The information and data stored in the cloud
- Devices that are allowed to connect to your cloud (cell phones, computers, and so on)
- The accounts and identities of the people, services, and devices within your organization
The cloud provider is always responsible for:

- The physical datacenter
- The physical network
- The physical hosts

# Cloud models
The three main cloud models are:
- Private : It’s a cloud (delivering IT services over the internet) that’s used by a single entity
- Public :  is built, controlled, and maintained by a third-party cloud provider
- Hybrid :  is a computing environment that uses both public and private clouds in an inter-connected environment
- Multi-cloud : use multiple public cloud providers(like AWS and Azure) and manage resources and secuirty in both enviroments

# Cloud benefits

- High availability and scalability(vertical, adding more CPU, more RAM to VMs and horizontal scalling, adding more servers)
- Realiability and predictability
- Security and governance
- Manageability in the cloud

# Cloud service types
# Infrastructure as a Service(IaaS)
Infrastructure as a service (IaaS) is the most flexible category of cloud services, as it provides you the maximum amount of control for your cloud resources. In an IaaS model, the cloud provider is responsible for maintaining the hardware, network connectivity (to the internet), and physical security. You’re responsible for everything else: operating system installation, configuration, and maintenance; network configuration; database and storage configuration; and so on. With IaaS, you’re essentially renting the hardware in a cloud datacenter, but what you do with that hardware is up to you.
# Platform as a Service(IaaS)
Platform as a service (PaaS) is a middle ground between renting space in a datacenter (infrastructure as a service) and paying for a complete and deployed solution (software as a service). In a PaaS environment, the cloud provider maintains the physical infrastructure, physical security, and connection to the internet. They also maintain the operating systems, middleware, development tools, and business intelligence services that make up a cloud solution. In a PaaS scenario, you don't have to worry about the licensing or patching for operating systems and databases.
# Software as a Service(IaaS)
Software as a service (SaaS) is the most complete cloud service model from a product perspective. With SaaS, you’re essentially renting or using a fully developed application

# Azure physical infrastructure
A region is a geographical area on the planet that contains at least one, but potentially multiple datacenters that are nearby and networked together with a low-latency network. like South Africa North region, Central US
Availability zones are physically separate datacenters within an Azure region
Azure regions are paired with another region within the same geography (such as US, Europe, or Asia) at least 300 miles away.

# Azure management infrastructure
A resource is the basic building block of Azure
Resource groups are simply groupings of resources(no nested resource groups)Resource groups provide a convenient way to group resources together. When you apply an action to a resource group, if deleted all the resorcess in it are also deleted.
subscriptions are a unit of management, billing, and scale. Similar to how resource groups are a way to logically organize resources, subscriptions allow you to logically organize your resource groups and facilitate billing
The final piece is the management group. Resources are gathered into resource groups, and resource groups are gathered into subscriptions. If you’re just starting in Azure that might seem like enough hierarchy to keep things organized. But imagine if you’re dealing with multiple applications, multiple development teams, in multiple geographies.

# Lab 1
# Create an Azure reource group
Steps
- Signin to the Azure portal
- Select create a resource > VM > Create 
- In basics tabs select subscription type(free student)
- Resource group(select the reource group created earlier)
- Enter name
- Authentication type :Password
- Set Username
- Set password
- Leave other settings as defaults 
# Create an Azure reource(Virtual machine)
Steps
- Signin to the Azure portal
- Select create a resource > VM > Create 
- In basics tabs select subscription type(free student)
- Resource group(select the reource group created earlier)
- Enter name
- Authentication type :Password
- Set Username
- Set password
- Leave other settings as defaults
# Clean up

It's a good idea at the end of a project to identify whether you still need the resources you created. Resources that you leave running can cost you money. You can delete resources individually or delete the resource group to delete the entire set of resources.