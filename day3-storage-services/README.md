## Azure storage services

Storage account provide a unique namespace for your azure storage data that is accessible from anywhere inthe world over HTTP/HTTPS
Storage account type
- Standard general purpose
- Premiun block blob
- Premium file share
- Premium page blobs

# Account endpoints
Combination of storage account name and the azure storage service endpoint form endpoint for your storage account 
Rules when naming 
- 3 - 24 characters(numbers and lowercase only)
- Unique in azure

# Azure storage redundancy
Azure always maintain multiple copies of data
Redundancy in primary region
- Locally Redundatant Storage: atleast 3 copies in a singe datacentre
- Zone Redundatant Storage: atleast 3 copies across Availability zones in the region

Redundancy in secondary region
- Geo-Redundatant Storage
- Geo-Zone Redundatant Storage

# Storage services
- Azure Blob - massively scalable object store for data
- Azure files - maanaged file share for cloud 
- Azure queues - messaging store
- Azure disks - for VMs
- Azure table - NoSQL table option for non relational data

# Benefits
- Durable and highly available
- Secure
- Scalable
- Managed services
- Accessible

# Access Tiers
- Hot access tier: frequently accessed data
- Cool access tier: infrequently accessed data(at least 30days)
- Cold access tier: infrequently accessed data(at least 90days)
- Archive access tier: infrequently accessed data(at least 180days)

# Data migartion
- Azure Migrate
- Azure Data Box

# Azure file movement
- AzCopy : CLI
- AZure storage explorer: GUI
- Azure file sync

# Lab
Create a storage account
- Sign in to the Azure portal at https://portal.azure.com
- Select Create a resource.
- Under Categories, select Storage.
- Under Storage account, select Create.
- Under Data storage create containers and upload files
- Select the blob(file) under properties tab copy URL(link to the file)