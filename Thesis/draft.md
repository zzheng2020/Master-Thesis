# High Availability in Lifecycle Management of Cloud-Native Network Functions: A Near-Zero Downtime Database Version Upgrade Prototype
* Content
  1. Introduction
     1. Overview
     2. Problem Statement
     3. Purpose
     4. Goals
     5. Research Methodology
     6. Delimitations
     7. Sustainability and Ethics
     8. Structure of the thesis
  2. Background
  3. System Design
  4. System Implementation
  5. Evaluation
  6. Conclusion

# 1. Introduction

This chapter …

## 1.1 Overview

In today's rapidly evolving technological landscape, an increasing number of companies are leveraging cloud services to improve various aspects of their businesses, such as controllability, flexibility, scalability, security, and resource utilisation. Ericsson, a leading telecommunications company, is among those that have integrated cloud services into their operations. In particular, Ericsson extensively employs Kubernetes for container orchestration and management, with a focus on the crucial subject of database lifecycle management.

## 1.2 Problem Statement

Within a clustered infrastructure, databases frequently encounter the need for version upgrades or downgrades to accommodate client demands. Conventionally, these modifications necessitate system downtime, which subsequently incapacitates the entire system and inhibits service provision to end-users throughout this interval. As the pursuit of high availability grows increasingly paramount for a multitude of enterprises, it becomes crucial to guarantee uninterrupted services for clients while concurrently bolstering the organisation's business development objectives.

## 1.3 Purpose

Given the importance of minimising downtime during database version changes, the aim of this project is to develop an efficient and streamlined approach. By reducing the potential for human error, we hope to increase the overall reliability and stability of the system during these transitions.

## 1.4 Goals

Our primary goals are as follows:

- Minimize the downtime required for database version changes, thereby ensuring uninterrupted service provision to customers.
- Automate the process of database version transitions to reduce the possibility of potential mistakes due to human factors.
- Enhance the overall system availability, contributing to improved customer satisfaction and the facilitation of business growth for Ericsson and its clients.
- Develop a robust and scalable solution that can be easily adapted to accommodate future technological advancements and changing business requirements.

By achieving these goals, we aim to create a more resilient and agile system that can better serve the needs of Ericsson and its customers in an increasingly competitive and dynamic business landscape.

## 1.5 Research Methodology

To achieve these goals, we propose employing a systematic research methodology that encompasses a combination of qualitative and quantitative research methods. This comprehensive approach will enable us to gather in-depth insights and develop a robust solution for minimising downtime during database version changes and automating the process. The research methodology comprises the following steps:

1. Literature Review: An extensive literature review will be conducted to explore existing research and solutions related to database lifecycle management, version changes, and high availability in cloud environments. This step will provide a solid understanding of the current state-of-the-art and help identify potential gaps and opportunities for improvement. The literature review will also provide insights into relevant methodologies, tools, and best practices that can be adopted or adapted for this project.
2. Requirement Analysis: Through interviews and surveys, we will engage with relevant stakeholders, including database administrators, system architects, and developers, to gather their perspectives and requirements for minimising downtime during database version changes and automating the process. This qualitative approach will help us understand the practical challenges faced by professionals in the field and ensure that the proposed solution aligns with their needs and expectations.
3. Design and Development: Based on the findings from the literature review and requirement analysis, we will design and develop a prototype solution for minimising downtime and automating database version changes. This solution will be guided by industry best practices and take into account the specific context and requirements of Ericsson and its customers.
4. Evaluation: Once the prototype solution is developed, we will conduct a series of quantitative experiments and performance tests to evaluate its effectiveness in minimising downtime and automating the process of database version changes. Key performance metrics, such as downtime duration, resource utilisation, and others will be measured and compared with pre-defined  goals. This quantitative approach will enable us to assess the efficacy of our solution and identify potential areas for improvement.
5. Iterative Refinement: Based on the results of the quantitative evaluation, we will iteratively refine and optimise the prototype solution to address any identified issues and enhance its performance. This process will be repeated until the solution meets the pre-defined benchmarks and goals.
6. Validation and Verification: Finally, the refined solution will be subjected to a rigorous validation and verification process to ensure its reliability, scalability, and adaptability in a real-world context. This step will involve further testing, as well as feedback from stakeholders and end-users, to confirm that the solution effectively addresses the identified challenges and meets the established goals.

By following this comprehensive research methodology, we aim to develop a robust, scalable, and adaptable solution that effectively minimises downtime during database version changes, automates the process, and ultimately enhances system availability and business growth for Ericsson and its customers.

## 1.6 Delimitations

This section outlines the delimitations and boundaries of the project, clarifying the scope and limitations to better focus the research effort and ensure a comprehensive understanding of the context. The following delimitations apply to the current project:

1. **Technological Stack**: The study will concentrate on Ericsson's specific technological stack, with a particular focus on Kubernetes for container orchestration and management. Consequently, the findings and proposed solution may not be directly applicable to other platforms or technological stacks without further adaptation.
2. **Database Version Changes**: The primary focus of this project is on database version upgrades and downgrades. Other forms of database maintenance, such as data migration or schema changes, are beyond the scope of this research and will not be addressed in the proposed solution.
3. **Ericsson's Business Context**: This research is conducted within the context of Ericsson's business operations and infrastructure, which may limit the generalisability of the findings and solution to other organisations. While the principles and techniques discussed may be applicable to a broader range of companies, additional research would be necessary to confirm their suitability for different business environments.
4. **Solution Scalability**: The proposed solution is designed to be scalable and adaptable to accommodate future technological advancements and changing business requirements. However, this research does not explore every possible future scenario or change in detail, and the solution may need further adjustments to address unforeseen challenges.
5. **Human Error**: The project aims to reduce the potential for human error during database version changes, but it does not claim to eliminate all human-related risks. Some degree of human involvement may still be necessary in certain aspects of the process, which could introduce potential mistakes.

By delineating these delimitations, we aim to create a focused and relevant research project that addresses the specific needs and challenges faced by Ericsson in their database lifecycle management. Despite the limitations outlined above, we believe that the insights gained and the solution developed will significantly contribute to improving Ericsson's operational efficiency, customer satisfaction, and overall business growth.

## 1.7 Sustainability and Ethics

This section discusses the sustainability and ethical considerations associated with the research project. The development of a solution for minimising downtime during database version changes has several implications in terms of environmental, social, and economic sustainability, as well as ethical concerns.

### 1.7.1 Environmental Sustainability

By minimising system downtime and increasing overall system efficiency, the proposed solution has the potential to reduce energy consumption and related environmental impacts. A more efficient system requires fewer resources to maintain, which can lead to a reduced carbon footprint for Ericsson and its clients. Furthermore, automating the database version transition process can contribute to reducing electronic waste, as it may decrease the need for hardware replacement due to fewer failures or malfunctions.

### 1.7.2 Social Sustainability

The proposed solution can have a positive impact on social sustainability by enhancing the overall system availability and reliability, leading to improved customer satisfaction. By ensuring uninterrupted service provision, customers will experience fewer disruptions in their daily lives and businesses, which may contribute to increased trust in the services provided by Ericsson and its clients.

### 1.7.3 Economic Sustainability

Streamlining the database version change process and minimising downtime can lead to significant cost savings for Ericsson and its clients. Reduced downtime translates into lower operational costs and fewer financial losses resulting from service disruptions. Moreover, a more efficient and automated process reduces the possibility of human error, which may decrease the need for expensive troubleshooting and problem resolution.

### 1.7.4 Ethical Considerations

The automation of the database version transition process may have implications for the job market and workforce, particularly for those involved in manual database management and maintenance tasks. As a result, it is essential to consider the potential displacement of workers and the need for retraining and upskilling to ensure that employees remain competitive in an increasingly automated industry.

Furthermore, the project must adhere to relevant data protection regulations and ensure that the proposed solution does not compromise the security and privacy of user data. The research and development process should take into account ethical principles such as transparency, accountability, and respect for user autonomy.

Therefore, this research project aims to address the challenges of minimising downtime during database version changes while considering the environmental, social, and economic sustainability aspects, as well as ethical concerns. By developing a robust, efficient, and scalable solution, we strive to contribute to a more sustainable and ethical technological landscape that benefits Ericsson, its clients, and society at large.

# 2. Background

```latex
\chapter{Background}
% \thispagestyle{fancy}
In this chapter, a detailed description about background of the degree project is presented together with related work. Discuss what is found useful and what is less useful. Use valid arguments. 

Explain what and how prior work / prior research will be applied on or used in the degree project /work (described in this thesis). Explain why and what is not used in the degree project and give valid reasons for rejecting the work/research.

Use references!

\section{Use headings to break the text}
Do not use subtitles after each other without text in between the sections.

\subsection{Related Work}
You should probably keep a heading about the related work here even though the entire chapter basically only contains related work.
```

This chapter aims to provide an overview of the contextual framework that underpins the Master Thesis Project.

## 2.1 Evolution of Software Architecture

Software architectures have undergone significant evolution over time. In today's software landscape, the main architectures can be broadly divided into two categories: monolithic applications and microservice architectures.

### 2.1.1 Monolithic Application

In traditional software development, a monolithic structure is widely used where it is characterised as a singular, self-contained system. Typically, such applications follow a layered architecture with a packaged approach to code management. The three-tier Model-View-Controller, or MVC, architecture is a commonly employed design pattern for monolithic applications, comprising presentation, business, and persistence layers \cite{deacon2009model}.

However, the monolithic architecture exhibits several limitations that hinder its scalability and maintainability.

* Version Control Challenges
  In a monolithic application, all code for a project is managed within a single code repository. When multiple developers work on the project, they each upload changes to the same repository. Each developer is responsible for one or more functional modules, and their commits are merged into the same repository. As a result, the repository may contain numerous branches and commits, making version control complex. This complexity often leads to conflicts in the merged code and can result in deployed code not containing the latest version for everyone involved.

* Maintainability Challenges

  Another limitation of monolithic architecture is its poor maintainability, particularly for large applications. An application with high maintainability enables developers to easily modify its code. However, when the amount of code within a project is substantial, developers can face challenges in navigating the codebase and keeping up with new developments. This can lead to delays and difficulty in implementing changes, thereby reducing maintainability.

* Resource Isolation Challenges

  In addition, monolithic architecture lack of effective resource isolation. Once a monolithic application is deployed, it exists in either a running or stopped state. In this context, resources are not adequately isolated, and any issues that occur can result in all customers using the application experiencing disruptions. This lack of resource isolation increases the risk of system crashes and can lead to negative consequences for end users.

### 2.1.2 Microservices

The microservices architecture is designed to decouple solutions by breaking down functionality into discrete services \cite{dragoni_microservices_2017}. Rather than using a monolithic structure, a microservices architecture employs several supporting microservices, which are developed, managed, and iterated on independently. This approach enables engineers to create applications that revolve around domain components, making product delivery easier by leveraging cloud architecture and platform-based deployment, management, and service capabilities. By using decentralised components, engineers can more efficiently manage the various services that comprise the application.

The microservices architecture provides several benefits over traditional monolithic architecture, including:

* Decoupled Services

  One of the primary benefits of microservices architecture is that there is no coupling between services. This allows for better scalability, as the microservice units of a microservice system have a strong ability to expand horizontally.

* Business Splitting

  Microservices are split by business, with solid service boundaries. This makes it easy to rewrite code for a business without having to understand the entire business.

* Independent Deployments

  Each service unit of a microservice is deployed independently, running in a separate process. This means that modifications and deployments of microservices have no impact on other services, improving system stability and reducing the risk of downtime.

## 2.2 Container

In the distant past, our approach to service deployment was to directly deploy them on hardware servers. However, this method presented numerous challenges when expanding services. To accommodate additional services, it was necessary to purchase additional servers and manually perform environment and service configurations, which were often error-prone and time-consuming, resulting in increased labor costs. As a consequence, service deployment and migration became highly inefficient. Initially, this method was suitable when the number of users and business volume on the Internet was limited. However, with the rapid expansion of businesses and the surge in the number of users, this approach to software service production was no longer viable. Deploying application services directly on the server presented three primary issues.

* Interference of services
  A server does not usually deploy just one service application, but multiple service applications. But as these services are all sharing server resources such as CPU, memory, hard disk and network IO in the server, then there is bound to be a situation where resources compete with each other and services affect each other.

* Low resource utilisation
  There are peaks and troughs in business, and when there is a trough, the resource utilisation of the server is much lower than when there is a peak in business. Therefore, the actual server resources are not utilised to their full potential during the business downturn.

* Difficulty in migration and expansion
  When the number of existing servers is not enough to cope with the fast-growing business, it is necessary to continuously expand the number of server instances. However, as the services are deployed directly in the servers, various dependency libraries, environment configurations and network configurations are required to migrate and expand the services, which is a complicated process and difficult to expand.

Therefore, container technology has emerged as a solution to the challenges presented by directly deploying services on hardware servers.

Containerization is a software deployment technique that consolidates the code of an application with all the necessary files and libraries essential for the application to operate on any infrastructure \cite{noauthor_what_nodate}. This deployment method offers a multitude of advantages, such as portability, scalability, fault tolerance, and agility.

Portability is a significant benefit of containerization that enables software developers to deploy applications across multiple platforms without the need to rewrite the code. This simplifies the development process, as only a single application build is required for deployment to multiple operating systems. For instance, developers can run the same container on both Linux and Windows operating systems, making it more efficient to upgrade legacy application code to modern versions.

Scalability is another advantage of containerization, as containers are lightweight software components that run efficiently. Virtual machines can start containerized applications more quickly, as they do not require a boot operating system. This enables software developers to add multiple containers for various applications on a single computer, with container clusters sharing computing resources from the same operating system without interfering with each other.

Fault tolerance is an essential feature of containerization, where software development teams can build fault-tolerant applications by using multiple containers to run microservices on the cloud. Containerized microservices operate in a separate user space, and if one container fails, it does not affect other containers, resulting in increased application resilience and availability.

Moreover, agility is an important aspect of containerization. Containerized applications operate in a separate computing environment, enabling software developers to troubleshoot and make changes to the application code without disrupting the operating system, hardware, or other application services. This accelerates software release cycles and enables quick updates using the container model.

## 2.3 ADP Plateform

Ericsson has developed the Application Development Platform (ADP) ecosystem to facilitate the development of cloud-native applications and services \cite{9758043}. This platform offers extensive technical, process, and organizational support to developers who create cloud network functions or cloud-native applications within Ericsson. The ecosystem is equipped with various tools, processes, and guidelines that enable developers to build top-quality applications that align with the business's needs.

The ADP ecosystem's architecture is based on modern microservices and container principles, which are widely utilized in contemporary software development. The microservices approach enables developers to deconstruct complex applications into smaller, more manageable components that can be developed and deployed independently. Containers offer a lightweight approach to bundling these components, making them more manageable and easier to move between different environments.

Furthermore, the ADP ecosystem provides critical support for the large-scale reuse of microservices across various applications within Ericsson. This feature helps developers reduce development time and costs by enabling them to create microservices once and then reuse them across multiple applications. The ADP Program also offers numerous resources, such as templates, examples, FAQs, and tutorials, to guide developers in utilizing the platform effectively.

The ADP ecosystem is an excellent platform for developers seeking to create cloud-native applications or cloud network functions within Ericsson. Its comprehensive support for technical, process and organizational aspects of software development makes it an ideal choice for teams looking to streamline their development processes while maintaining high levels of quality and efficiency.

## 2.4 CAP Theorem

In a distributed system, it is possible to achieve at most two out of the following three properties: Consistency, Availability, and Partition Tolerance.

### 2.4.1 Consistency

Consistency refers to the condition where "all nodes see the same data at the same time." That is, after an update operation is successfully completed and returned to the client, all nodes have the same data at the same time. Thus, consistency concerns data uniformity in a distributed system.

Consistency can be viewed from two different perspectives: the client-side and the server-side. From the client-side, consistency mainly deals with how updated data is accessed during concurrent operations. From the server-side, it concerns how updates are replicated and distributed throughout the system to ensure eventual consistency.

Consistency issues arise due to concurrent read and write operations; therefore, it is crucial to consider concurrent scenarios when understanding consistency problems.

From the client-side perspective, different consistency policies dictate how updated data is accessed by different processes during concurrent operations.

Here are three Consistency Policies:

1. Strong Consistency: In the context of relational databases, strong consistency requires that any updated data be immediately visible to all subsequent accesses.

2. Weak Consistency: Weak consistency tolerates situations where some or all subsequent accesses may not immediately see the updated data.

3. Eventual Consistency: Eventual consistency ensures that the updated data becomes visible after a certain period.

### 2.4.2 Availability

Availability refers to the condition where "reads and writes always succeed," meaning that the service is consistently accessible and provides normal response times.

In a distributed system with high availability, every non-failing node must respond to every request. Therefore, availability is generally measured by assessing system downtime. When describing the availability of a system, for example, we may say that the e-commerce platform Taobao achieves "five nines" availability, indicating a 99.999% uptime. This means that the total annual downtime does not exceed (1-0.99999) * 365 * 24 * 60 = 5.256 minutes, which is an exceptionally high standard.

Good availability signifies that the system can effectively serve its users without experiencing issues such as operation failures or access timeouts, which could lead to poor user experiences. In a distributed system, various components like load balancers, web servers, application code, and database servers contribute to the overall design. The stability of any single component can directly impact the system's availability.

### 2.4.3 Partition Tolerance

Partition tolerance refers to the ability of "the system to continue operating despite arbitrary message loss or failure of part of the system." In other words, a distributed system can still provide services that meet consistency and availability requirements, even when encountering node or network partition failures.

Partition tolerance is closely related to scalability. In distributed applications, system malfunctions may arise due to various distributed factors. Good partition tolerance ensures that the application appears to function as a coherent whole, even though it is a distributed system. For instance, if one or several machines in a distributed system fail, the remaining machines can still operate normally and meet system requirements. Alternatively, if there is a network anomaly that separates the distributed system into independent parts, each part can still maintain the operation of the distributed system. This demonstrates effective partition tolerance.

In simpler terms, if the system can still function normally under circumstances of network disruption or message loss, it is considered to possess good partition tolerance.

## 2.5 DevOps and Continuous Delivery

DevOps and continuous delivery are interrelated practices that work in tandem to foster a culture of collaboration and automation to optimize software delivery. The goal of these practices is to enhance software delivery by facilitating faster, more reliable, and high-quality releases through automation and collaborative efforts.

DevOps, a portmanteau of "development" and "operations," is an innovative approach that aims to integrate software development and IT operations seamlessly. This interdisciplinary methodology is designed to bridge the gap between these traditionally siloed domains, fostering collaboration, communication, and efficiency. By leveraging a suite of tools, practices, and cultural shifts, DevOps aims to enhance the software development lifecycle (SDLC) through continuous delivery, automated testing, and rapid iteration. As a result, organizations employing DevOps principles can improve the quality, reliability, and speed of their software deployments. Moreover, this holistic approach promotes a culture of shared responsibility, enabling teams to respond more effectively to changing requirements and mitigating risks associated with traditional development methodologies. In essence, DevOps represents a paradigm shift in software engineering, with the potential to reshape organizational structures and drive significant advancements in the field.

Continuous Delivery (CD) is a software engineering approach that emphasizes the automation and streamlining of the software release process, thereby ensuring rapid and reliable deployment of new features and bug fixes. As an integral component of the DevOps methodology, CD fosters a culture of iterative improvement, facilitating seamless integration of code changes into a shared repository, and subsequently delivering these changes to production environments with minimal human intervention.

CD relies on a robust suite of tools and practices, including version control, automated build processes, testing frameworks, and deployment mechanisms, which collectively contribute to reduced lead times and increased deployment frequency. By promoting a systematic, reproducible, and auditable release process, CD mitigates risks associated with manual deployments and fosters greater collaboration among development and operations teams. Furthermore, this approach enables organizations to respond swiftly to changing market demands, enhance customer satisfaction, and maintain a competitive edge. Moreover, Continuous Delivery represents a transformative paradigm within the software engineering landscape, driving innovation and efficiency in an increasingly complex and rapidly evolving technological landscape.

## 2.6 Kubernetes

As previously discussed, Docker proves to be a powerful tool for the development of microservices systems. However, in the context of clusters comprising tens of thousands of containers, numerous challenges arise in terms of scheduling, managing, and dispatching these resources. Consequently, an efficient and flexible management system is required to address these concerns. Kubernetes was developed as a solution to these issues, with its name derived from the Greek word for "helmsman" or "pilot." The platform builds on Google's extensive experience in managing large-scale production workloads and incorporates the best practices and ideas from the community.

Kubernetes is a portable, scalable, open-source platform designed for the management of containerized workloads and services. It enables declarative configuration and automation, while boasting a large, rapidly growing ecosystem of services, support, and tools. Utilizing Kubernetes allows for rapid application deployment, efficient scaling, seamless integration of new application features, resource conservation, and optimization of hardware resource utilization.

Furthermore, Kubernetes offers several key features, including:

1. Portability: Support for public, private, hybrid, and multi-cloud environments.
2. Extensibility: A modular architecture that allows for plug-ins, mounting, and combination of components.
3. Automation: Facilitation of automatic deployment, restart, replication, and scaling/extension.

### 2.6.1 Kubernetes Architecture

To provide a concise overview, the architecture of Kubernetes consists of a Master managing a group of Nodes.

The Master possesses the following components:

1. API server: The gateway for all command requests in Kubernetes.
2. Scheduler: Utilizes a scheduling algorithm to allocate requested resources to a Node.
3. Controller: Manages Kubernetes resource objects.
4. etcd: Stores the resource objects.

In contrast, Nodes encompass the following elements:

1. Kubelet: Present on each Node, it executes resource manipulation instructions.
2. Kube-proxy: Manages load balancing between services.
3. Pod: Represents the fundamental unit (smallest manageable unit) within Kubernetes, which houses the container. Kubernetes manages Pods rather than individual containers.
4. Docker: Provides the base environment for container execution, functioning as the container engine.
5. Fluentd: Offers log collection services.

This structure enables Kubernetes to efficiently manage containerized workloads and services while maintaining a high level of automation and flexibility.

### 2.6.2 Kubernetes Core Components

Kubernetes encompasses a variety of fundamental components, including Pods, Kubelet, ReplicaSet, and Deployment, which together form the basis for efficient container management and orchestration within the platform.

Kubernetes serves as a container management system, though it does not directly interact with containers. Instead, it operates on the smallest unit called a Pod, which indirectly manages containers. A Master node corresponds to a group of Node nodes, with the Master node responsible for scheduling, network management, controller, and resource object storage. Containers are housed within the Node node, which in turn, stores them inside the Pod. Pods can encapsulate single or multiple containers.

Kubelet is tasked with local Pod maintenance, while kube-proxy manages load balancing between multiple Pods. A Pod is a container that encapsulates other containers created by Docker, functioning as a virtualized group. Pods act as standalone hosts and can house one or more containers. Each Pod possesses its own IP address and hostname, providing an independent sandbox environment.

Pods are typically employed for managing groups of related services during service deployment. They can deploy a single service or a group of related services, which are sets of services connected through a call path of a chain call. Kubernetes excels at managing web service clustering through the control of the number of Pods, while the underlying network of Pods and data storage is facilitated by the creation of Pause containers prior to the establishment of internal containers. Service containers can access localhost, offering high-performance access to local services.

ReplicaSet controllers maintain the desired number of Pod replicas or "service clusters," ensuring consistent replica numbers. In case a Pod service fails, the ReplicaSet controller immediately recreates a new Pod. Replica Controllers utilize tag selectors to maintain a group of related services. In newer versions of Kubernetes, ReplicaSet is recommended over the deprecated ReplicationController.

Deployment objects represent the service deployment architecture model, enabling rolling updates. While ReplicaSet controllers control the number of Pod replicas, they do not support rolling updates. Deployment objects, on the other hand, support rolling updates and are often used in conjunction with ReplicaSets. Deployment manages ReplicaSets, with the latter reestablishing new ReplicaSets and creating new Pods to facilitate project version updates when needed.

### 2.6.3 Stateful vs. Stateless

Stateful and Stateless resources are critical concepts in Kubernetes, particularly when addressing containerized deployment challenges. When deploying PostgreSQL within containers, data loss becomes a significant concern due to the lifecycle of both containers and Pods. Kubernetes does not utilize Deployment for managing stateful services; instead, it employs Deployment for stateless services and StatefulSet for stateful services deployment.

Stateful services are characterized by real-time data storage requirements. In a stateful service cluster, when a service is temporarily removed and later rejoined to the network, the unavailability of the cluster network may present issues. On the other hand, stateless services do not necessitate real-time data storage. Consequently, in a stateless service cluster, temporary removal and reintegration of a service into the network will not impact the clustered service.

StatefulSet is employed to address the challenges associated with using containerized deployment for stateful services. This deployment model is specifically designed for stateful services, ensuring that the Hostname remains unaltered after Pod recreation. This stability enables the Pod to associate data through the Hostname, ultimately improving the management and resilience of stateful services within Kubernetes.

### 2.6.4 Kubernetes Operator

The Kubernetes Operator is a concept designed to encapsulate the knowledge and objectives of an operations professional responsible for managing a service or a set of services. These professionals possess an in-depth understanding of system operations, deployment, and problem resolution. Automation is a preferred method for handling repetitive tasks in the context of Kubernetes workloads, and the Operator pattern serves to embody the task automation code written outside the functionality provided by Kubernetes itself \cite{dobies2020kubernetes}.

Operators function by extending the Kubernetes control plane and API. They introduce an endpoint (referred to as a custom resource or CR) to the Kubernetes API, which also incorporates a control plane component that monitors and maintains new resource types. An Operator comprises a set of Controllers, each listening to specific Kubernetes resources. These Controllers can implement a reconciliation loop, with each Controller being responsible for monitoring a designated resource and triggering reconciliation upon creation, update, or deletion of the monitored resource.

Several open-source projects facilitate the creation of Kubernetes Operators, such as Operator SDK, Kubebuilder, KUDO, and Metacontroller. These tools streamline the development and management of Operators within the Kubernetes ecosystem.

Kubernetes Controllers, on the other hand, are non-terminating loops that regulate the state of the system, striving to align the current state with the desired state as closely as possible through a process known as the Reconciliation loop. In the context of Kubernetes, a set of built-in Controllers operates within the controller-manager located in the master node. These Controllers work in tandem with Operators to ensure efficient and accurate management of resources in a Kubernetes environment.

## 2.6 Flux

Flux is a specialized tool designed to maintain synchronization between Kubernetes clusters and configuration sources, such as Git repositories. This synchronization enables the automation of updates to the configuration when new code deployments arise. Developed with a focus on leveraging Kubernetes' API extension system, Flux seamlessly integrates with core components of the Kubernetes ecosystem, including Prometheus.

One of the key strengths of Flux is its support for multi-tenancy, allowing for the efficient management of multiple tenants within a single cluster. Additionally, Flux facilitates the synchronization of an arbitrary number of Git repositories, providing flexibility and adaptability in managing diverse sources of configuration. This capability ensures that the Kubernetes clusters remain up-to-date with the latest configurations, enhancing the efficiency and reliability of container management and orchestration.

# 3. System Design

This chapter …

## 3.1 Goals in System Design

This section outlines the primary objectives we aim to accomplish in the design of a system that addresses the challenges associated with database version changes. These goals have been established to guide the development of a solution that effectively minimises downtime, automates processes, and remains transparent to clients.

The main goals in the system design are as follows:

1. **Perform Database Version Change**: Enable efficient and seamless transitions between different database versions, catering to the evolving needs of clients and businesses.
2. **Minimise Downtime**: Reduce the duration of system downtime during database version changes to ensure uninterrupted service provision and maintain high availability for users.
3. **Automate the Procedure**: Implement automation to minimize manual intervention in the database version change process, decreasing the potential for human error and increasing overall efficiency and reliability.
4. **Transparent to Clients**: Ensure that the process of database version changes remains transparent to clients, allowing them to continue using the services without disruption or any discernible difference in performance during the transition.

By focusing on these objectives in the system design, we aim to develop a comprehensive solution that addresses the limitations of conventional strategies while offering enhanced efficiency, reliability, and transparency.

## 3.2 Conventional Strategy

This section describes the conventional strategy for implementing a change in the database version and discusses its limitations with respect to the need for high availability in modern business environments.

The traditional implementation scheme for changing the database version consists of the following steps:

1. Suspend the service, ceasing to receive user requests.
2. Create a backup of the database that requires an upgrade.
3. Establish a target version database within the cluster, and migrate the data from the backup created in step 2 to this target version database.
4. Verify the consistency of the data in the target version database with the data in the original database.
5. If the verification in step 4 is successful, resume the service and begin receiving user requests once more.

While this approach is effective for managing database version changes, it necessitates a certain amount of system downtime. The duration of this downtime is contingent upon the speed of data migration and validation; the faster these processes, the less the service will be impacted.

In today's business landscape, however, there is an increasing demand for high availability in conjunction with system stability. Table 1 illustrates common requirements for high availability. For instance, to attain the "five nines" target (i.e., 99.999\% high availability), services must experience minimal downtime. Evidently, the traditional implementation strategy is inadequate for meeting such stringent requirements. Consequently, the need for an alternative solution that can minimise downtime and maintain high availability during database version changes becomes paramount.

\todo[inline]{//TODO: Insert Table one here.}

## 3.3 Blue-Green Deployment

In order to address the challenges associated with database version changes and achieve the goals outlined in the system design, the blue-green deployment strategy emerges as a viable approach. This section provides a professional description of the blue-green deployment strategy and explains how it contributes to the objectives of minimising downtime, automating processes, and maintaining transparency for clients.

The blue-green deployment strategy involves creating two separate environments, referred to as "blue" and "green," which run concurrently. Each environment hosts a different version of the database, allowing for efficient and seamless transitions between versions with minimal downtime. The key advantage of this approach is that it enables continuous operation of the system while the database version is changed, thus addressing the goal of minimising downtime.

There are four main step in Blue-Green Deployment strategy:

1. **Preparation**: Initially, both blue and green environments are set up, with the blue environment hosting the current database version and the green environment hosting the target version. During this stage, the green environment is prepared and tested to ensure its readiness for the database version change.
2. **Data Synchronization**: Data is synchronized between the blue and green environments to ensure consistency. This process is automated to reduce manual intervention, contributing to the goal of automating the procedure and increasing overall efficiency and reliability.
3. **Switching**: Once the data synchronization is complete, the system switches from the blue environment to the green environment. This transition is seamless and occurs without any noticeable interruption in service provision, addressing the goal of maintaining transparency for clients.
4. **Monitoring**: After the switch, the system is closely monitored to ensure that the new database version is functioning correctly and that there are no performance issues or inconsistencies. If any issues are detected, the system can quickly revert to the blue environment, ensuring minimal disruption to clients.

The blue-green deployment strategy enables the efficient management of database version changes while minimising downtime, automating processes, and maintaining transparency for clients. By incorporating this strategy in the system design, we can develop a comprehensive solution that addresses the limitations of conventional strategies and better supports the needs of Ericsson and its clients in a competitive and dynamic business landscape.

### 3.3.1 Advantages of Blue-Green Deployment Strategy

The blue-green deployment strategy offers several advantages when implemented to address the challenges of database version changes. The followings outline the key benefits associated with this approach.

1. Reduced Downtime: By maintaining two separate environments and enabling a seamless switch between them, the blue-green deployment strategy significantly reduces the downtime associated with database version changes. This ensures that clients experience minimal disruption in service provision, addressing the primary goal of minimizing downtime.
2. Improved Reliability: The blue-green deployment strategy allows for thorough testing of the target version environment before transitioning to it. This ensures that potential issues can be identified and resolved before the switch, contributing to the overall reliability of the system during the database version change process.
3. Enhanced Flexibility: The ability to quickly switch between the blue and green environments provides flexibility in handling any issues that may arise during the transition. In the event of a problem, the system can easily revert to the previous version, ensuring minimal impact on service provision.
4. Risk Mitigation: The blue-green deployment strategy mitigates the risk associated with database version changes by enabling thorough testing and data synchronisation before the transition. This helps to prevent potential data loss or corruption during the process.
5. Increased Agility: The ability to quickly and efficiently change database versions using the blue-green deployment strategy enables organisations to adapt more rapidly to evolving client demands and market conditions. This increased agility can provide a competitive advantage in a dynamic business landscape.

Therefore, the blue-green deployment strategy offers a robust and effective solution for managing database version changes while addressing the key goals of minimising downtime, automating processes, and maintaining transparency for clients. By leveraging this approach in the system design, we can better support the needs of Ericsson and its clients in an increasingly competitive and dynamic business environment.

## 3.4 Synchronisation between PostgreSQL Databases

In the context of the blue-green deployment strategy, synchronisation between the two PostgreSQL databases is crucial for ensuring data consistency and maintaining transparency for clients during the database version change process. This section introduces logical replication as the primary method for achieving synchronisation between the two PostgreSQL databases, highlighting its advantages and relevance to the system design goals.

### 3.4.1 Logical Replication

Logical replication is a feature in PostgreSQL that enables the selective replication of data changes from one database to another. It facilitates real-time synchronisation between the two databases, ensuring that any changes made in the source database are instantly propagated to the target database. This approach is particularly well-suited for the blue-green deployment strategy, as it allows for seamless transitions between different database versions with minimal downtime and disruption to clients.

### 3.4.2 Advantages of Logical Replication

Logical replication offers several advantages that align with the system design goals, including:

- **Minimised Downtime**: By enabling real-time synchronisation between the two databases, logical replication allows for seamless transitions between different database versions without significant downtime, ensuring high availability for users.
- **Automated Procedure**: The replication process is largely automated, reducing manual intervention and the potential for human error, thereby increasing overall efficiency and reliability.
- **Transparent to Clients**: As the data synchronisation occurs in real-time, clients can continue using the services without disruption or any discernible difference in performance during the transition.

By incorporating logical replication in the synchronisation process, we can effectively address the challenges associated with database version changes, and achieve the system design goals of minimising downtime, automating processes, and maintaining transparency for clients.

## 3.5 Monitoring Synchronisation Progress

An essential aspect of ensuring a smooth database version change process is the ability to monitor the progress of synchronisation between the master and follower nodes. In PostgreSQL, the `pg_stat_replication` view provides information on the replication status, which can be utilised to determine when the synchronisation process is complete.

### 3.5.1 Key Replication Metrics

The following metrics provided by the `pg_stat_replication` view are crucial to understanding the synchronisation progress:

1. **sent_lsn**: This metric represents the latest Write-Ahead Log (WAL) position sent by the master node to the follower node.
2. **pg_current_wal_flush_lsn**: This function returns the latest WAL position that has been flushed to disk on the master node.
3. **pg_wal_lsn_diff(lsn1, lsn2)**: This function calculates the byte difference between two WAL positions.

### 3.5.2 Determining Synchronisation Completion

To determine when the synchronisation process is complete, we can compare the `sent_lsn` value with the `pg_current_wal_flush_lsn` value. When these two values match, it indicates that the latest WAL position sent by the master node has been flushed to disk, and the synchronisation process is complete.

Additionally, the `pg_wal_lsn_diff` function can be used to monitor the remaining byte difference between the master and follower nodes' WAL positions. This information can be useful in estimating the remaining synchronisation time, allowing for more accurate predictions and adjustments to the system.

By leveraging the information provided by the `pg_stat_replication` view and the associated functions, we can effectively monitor the synchronisation progress between the master and follower PostgreSQL databases. This capability aligns with our system design goals, ensuring minimal downtime, automated processes, and transparent operation for clients during database version changes. As a result, we can better support the needs of Ericsson and its clients in a competitive and dynamic business landscape.

## 3.6 Integrating Blue-Green Deployment Strategy with Kubernetes

As Ericsson extensively employs Kubernetes for container orchestration and management, it is essential to integrate the blue-green deployment strategy with Kubernetes to ensure seamless database version transitions. This section discusses the process of integrating the blue-green deployment strategy within a Kubernetes-based infrastructure and its benefits in the context of database lifecycle management.

Kubernetes is an open-source container orchestration platform that automates the deployment, scaling, and management of containerised applications. By integrating the blue-green deployment strategy with Kubernetes, we can leverage the platform's powerful features to manage the database version changes more efficiently.

Integrating the blue-green deployment strategy with Kubernetes offers several benefits, including:

- **Scalability**: Kubernetes enables easy horizontal scaling of the database instances, ensuring that the system can accommodate the evolving needs of clients and businesses.
- **Resource Utilization**: Containerization and the efficient use of resources within Kubernetes ensure optimal resource utilization, ultimately reducing costs and improving performance.
- **Ease of Management**: Kubernetes provides a unified platform for managing the entire database lifecycle, simplifying the process of implementing and maintaining the blue-green deployment strategy.

- **Fault Tolerance**: By leveraging the inherent fault tolerance and self-healing capabilities of Kubernetes, the integrated blue-green deployment strategy can effectively handle failures and maintain system stability during database version changes.
- **Monitoring and Observability**: Kubernetes offers built-in monitoring and observability tools, enabling real-time insights into the performance and health of the database instances during the transition process. This facilitates prompt identification and resolution of issues, ensuring a smooth and reliable transition.

By integrating the blue-green deployment strategy with Kubernetes, we can effectively address the challenges associated with database version changes while harnessing the powerful features of Kubernetes for container orchestration and management. This integration contributes to the achievement of the system design goals, including minimising downtime, automating processes, and maintaining transparency for clients, ultimately fostering a more resilient and agile system that can better serve the needs of Ericsson and its customers in a competitive and dynamic business landscape.

### 3.6.1 Achieving the Blue-Green Deployment with Kubernetes Operator

Kubernetes Operator is a framework for building Kubernetes-native applications, providing domain-specific knowledge to automate complex tasks. By integrating the blue-green deployment strategy with Kubernetes Operator, we can take advantage of its capabilities to manage the database version changes effectively, offering the following benefits:

1. **Automation**: Kubernetes Operator enables automation of complex tasks associated with the blue-green deployment process, such as data synchronization and traffic switching, thus reducing manual intervention and potential human errors.
2. **Extensibility**: Kubernetes Operator allows for the creation of custom resources that define the desired state of the system, making it adaptable to different database systems and future technological advancements.
3. **Ease of Management**: Kubernetes Operator provides a unified platform for managing the entire database lifecycle, simplifying the implementation and maintenance of the blue-green deployment strategy.

### 3.6.2 Custom Resource Definition (CRD)

Custom Resource Definitions (CRDs) are a crucial component of Kubernetes that extend its API, enabling users to define, create, and manage new custom resources. In the context of the blue-green deployment strategy for database version changes, we use a CRD specifically designed for this purpose.

To design the CRD for the blue-green deployment strategy, we define a schema that specifies the structure and properties of the custom resource. The schema includes the following fields:

1. **apiVersion**: The version of the API used to create the CRD.
2. **kind**: The type of the Kubernetes object, which in this case, is a custom resource related to blue-green deployment.
3. **metadata**: Contains metadata about the custom resource, such as name, namespace, labels, and annotations.
4. **spec**: The specification of the custom resource.

By utilising the CRD, we can create and manage custom resources for the blue-green deployment strategy within the Kubernetes ecosystem. This approach allows us to achieve the system design goals, including minimising downtime, automating processes, and maintaining transparency for clients, ultimately fostering a more resilient and agile system that can better serve the needs of Ericsson and its customers in a competitive and dynamic business landscape.

# 4. System Implementation

The system implementation chapter focuses on the code and logic implemented in the `PgUpgradeReconciler` to achieve the desired functionality. This chapter explains the code implementation and the processes involved in reconciling a PgUpgrade object to create and manage PostgreSQL upgrades.

## 4.1 Achieving Logical Replication between PostgreSQL Databases

Logical replication in PostgreSQL enables the efficient and seamless replication of data changes from a source (master) node to a target (follower) node. In the context of our system design, we leverage Kubernetes Operator to automate the configuration and management of logical replication between PostgreSQL databases.

### 4.1.1 Configuring the Master Node

To set up logical replication on the master node, Kubernetes Operator automates the following tasks:

1. **Configure PostgreSQL settings**: The Operator updates the `postgresql.conf` file to enable logical replication by setting the `wal_level` to `logical`.
2. **Create a publication**: The Operator uses the SQL command `CREATE PUBLICATION` to create a new publication that includes all tables.
3. **Export the schema**: The Operator employs the `pg_dump` command to export the schema of the database to an SQL file.

### 4.1.2 Configuring the Follower Node

Kubernetes Operator automates the synchronisation of the schema and the creation of a subscription on the follower node by performing the following tasks:

1. **Synchronise the schema**: The Operator imports the schema from the master node using the `psql` command.
2. **Create the subscription**: The Operator uses the SQL command `CREATE SUBSCRIPTION` to create a new subscription that connects to the master node's publication.

By leveraging Kubernetes Operator, we can successfully establish and automate logical replication between the master and follower PostgreSQL databases. This approach aligns with our system design goals, ensuring minimal downtime, automated processes, and transparent operation for clients during database version changes.

## 4.1 `PgUpgradeReconciler` Structure

The `PgUpgradeReconciler` struct consists of the following fields:

1. `client.Client`: An instance of the client, which is responsible for performing CRUD operations on Kubernetes objects.
2. `Scheme`: A runtime.Scheme object that enables conversion between Go structs and GroupVersionKinds.
3. `Log`: An instance of the logr.Logger to log messages at different verbosity levels.

## 4.2 Kubernetes RBAC

Kubebuilder annotations define the necessary RBAC permissions for the controller to manage resources:

1. `pgupgrades`: Custom resource for managing PostgreSQL upgrades.
2. `pgupgrades/status`: Subresource for managing the status of the PgUpgrade objects.
3. `pgupgrades/finalizers`: Subresource for managing finalizers on PgUpgrade objects.
4. `pods`, `deployments`, `services`, and `configmaps`: Standard Kubernetes resources required for creating, managing, and deleting related resources.

## 4.3 Reconciliation Process

The `Reconcile` function is the core of the controller and handles the reconciliation process. It starts by retrieving the PgUpgrade instance and checks if it exists. If not found, the reconcile function returns a nil error, allowing the controller to continue with other instances.

### 4.3.1 Controller Setup

The `PgUpgradeReconciler` struct is used to define the controller for the PgUpgrade system. The `SetupWithManager` function is responsible for setting up the controller with the Manager. It associates the controller with the `pgupgradev1.PgUpgrade` resource and returns the result.

### 4.3.1 Deployment Creation

The function checks if the PgUpgrade deployment already exists. If not, it creates a new deployment using the `deploymentForPgUpgrade` function and logs the creation process. If the deployment already exists, it ensures the deployment image is the same as specified in the PgUpgrade object. If there is a difference, it updates the deployment with the new image.

### 4.3.2 Service Creation

The function checks if the PgUpgrade service already exists. If not, it creates a new service using the `serviceForPgUpgrade` function and logs the creation process. If the service already exists, it moves on to the next step.

### 4.3.3 Schema Synchronisation

The `syncSchema` function executes the schema synchronisation command on the target PostgreSQL instance. It does so by connecting to the target pod, discovering its IP address, and running the schema sync command through the `remotecommand` package.

### 4.3.4 Subscription Creation

The `createSubscriptions` function establishes the logical replication subscription between the old and new PostgreSQL instances. It connects to the new PostgreSQL instance and runs the `create subscription` command with the necessary parameters.

### 4.3.5 Nginx Proxy Change

The `changeNginxProxyPass` function is responsible for updating the Nginx configuration stored in a ConfigMap. It takes the context, request, PgUpgrade resource, and deployment as input arguments. The function first retrieves the ConfigMap containing the Nginx configuration and checks if it contains the `nginx.conf` key. If the key is present, the function replaces the old database host value with the new follower IP in the `proxy_pass` directive. Then it updates the ConfigMap in the cluster.

### 4.3.6 Resource Deletion

The `deleteResource` function is responsible for deleting the specified resources (deployments) from the cluster. It takes the context, PgUpgrade resource, and deployment as input arguments. The function first creates a new Kubernetes clientset and then iterates over the resource names specified in the `KillDeployments` field of the PgUpgrade resource. For each resource, it deletes the deployment in the specified namespace and logs the deletion operation.

### 4.3.8 Helper Functions

Several helper functions are implemented to support the primary functions of the system. These include:

- `labelsForPgUpgrade`: Generates labels for a PgUpgrade resource.
- `getFollowerIP`: Be responsible for retrieving the follower's IP address from the cluster. It takes the context, PgUpgrade resource, and deployment as input arguments. The function first creates a new Kubernetes clientset and then retrieves the pod name for the specified `OldDBLabel`. After obtaining the pod name, it retrieves the pod IP address and logs it.

The System Implementation chapter describes the key components of the PgUpgrade controller and the methods for creating resources, syncing schemas, creating subscriptions, deleting resources, and updating the Nginx configuration. These components and functions work together to automate the PostgreSQL upgrade process and minimise downtime during the upgrade.

# 5. Prototype Outcomes

The prototype developed in this study aimed to address the challenges of database version migration and traffic switching, ensuring minimal disruption to users. The outcomes of the prototype are discussed below.

## 5.1 Continuous Read/Write Access

During the initial phase, users were able to send read/write requests to the database cluster via the exposed port on nginx. This allowed continuous access to the database even when the migration to the new version began. It ensured that users experienced no downtime or disruption in service during the migration process.

## 5.2 Database Version Migration and Synchronisation

In the cluster, a new version of the database was created and synchronised with the old version using logical replication. This approach facilitated a smooth transfer of data between the two versions while maintaining data consistency. By monitoring the information in pg_stat_replication, the prototype was able to automatically switch traffic from the old version to the new version at the appropriate time.

## 5.3 Transparent Traffic Switching

Upon completion of the migration process, the user's read/write requests were successfully switched to the new version of the database. The entire operation of database version migration and traffic switching within the cluster was transparent to users, allowing them to continue using the service without any awareness of the underlying changes.

The prototype outcomes demonstrate that the proposed solution is effective in addressing the challenges of database version migration and traffic switching. The seamless experience provided to users during the migration process underscores the practicality and robustness of the approach. Further research and development can focus on refining the prototype for broader applications and testing in various database environments.

# 6. Conclusions

This thesis aimed to design a system addressing the challenges associated with database version changes, with specific goals focused on minimising downtime, automating processes, and maintaining transparency to clients. The primary objectives in the system design were to:

1. Perform Database Version Change
2. Minimise Downtime
3. Automate the Procedure
4. Remain Transparent to Clients

The prototype presented in this thesis successfully demonstrated that the proposed solution could effectively meet the established goals. By providing continuous read/write access, seamless version migration and synchronisation, and transparent traffic switching, the system ensured uninterrupted service provision and minimal disruption to users.

## 6.1 Achievement

The achievements of this project in relation to the stated goals are as follows:

1. **Perform Database Version Change**: The prototype effectively managed the transition between different database versions by leveraging logical replication, ensuring data consistency and seamless migration.
2. **Minimise Downtime**: The system achieved minimal downtime by allowing users to send read/write requests during the migration process, ensuring uninterrupted access to the database.
3. **Automate the Procedure**: The prototype automated traffic switching through monitoring pg_stat_replication data, reducing the need for manual intervention and increasing overall efficiency and reliability.
4. **Transparent to Clients**: The entire process of database version changes remained transparent to clients, allowing them to continue using the services without any discernible difference in performance during the transition.

##  6.2 Future Work

While the prototype successfully demonstrated the effectiveness of the proposed solution, future work can be directed towards refining and extending the system for broader applications. Potential areas for future research and development include:

1. Testing the system in various database environments to assess its adaptability and compatibility with different database management systems.
2. Evaluating the scalability of the system in handling larger and more complex migrations, including those involving multiple database clusters and various data types.
3. Investigating potential improvements in the automation processes and monitoring tools to further optimise the efficiency and reliability of the system.

In conclusion, this thesis presents a successful prototype that addresses the challenges associated with database version changes. The developed system effectively minimises downtime, automates critical processes, and remains transparent to clients, demonstrating the practicality and robustness of the proposed solution. With further research and development, the system can be refined for broader applications, contributing to the evolving landscape of database management and service provision.

# Reference

```latex
@article{deacon2009model,
  title={Model-view-controller (mvc) architecture},
  author={Deacon, John},
  journal={Online][Citado em: 10 de mar{\c{c}}o de 2006.] http://www.jdl.co.uk/briefings/MVC.pdf},
  volume={28},
  year={2009}
}

@incollection{dragoni_microservices_2017,
	location = {Cham},
	title = {Microservices: Yesterday, Today, and Tomorrow},
	isbn = {978-3-319-67425-4},
	url = {https://doi.org/10.1007/978-3-319-67425-4_12},
	abstract = {Microservices is an architectural style inspired by service-oriented computing that has recently started gaining popularity. Before presenting the current state of the art in the field, this chapter reviews the history of software architecture, the reasons that led to the diffusion of objects and services first, and microservices later. Finally, open problems and future challenges are introduced. This survey primarily addresses newcomers to the discipline, while offering an academic viewpoint on the topic. In addition, we investigate some practical issues and point out a few potential solutions.},
	pages = {195--216},
	booktitle = {Present and Ulterior Software Engineering},
	publisher = {Springer International Publishing},
	author = {Dragoni, Nicola and Giallorenzo, Saverio and Lafuente, Alberto Lluch and Mazzara, Manuel and Montesi, Fabrizio and Mustafin, Ruslan and Safina, Larisa},
	editor = {Mazzara, Manuel and Meyer, Bertrand},
	date = {2017},
	doi = {10.1007/978-3-319-67425-4_12},
}


@misc{noauthor_what_nodate,
	title = {What is {Containerization}? - {Containerization} {Explained} - {AWS}},
	shorttitle = {What is {Containerization}?},
	url = {https://aws.amazon.com/what-is/containerization/},
	abstract = {Find out what is Containerization and how to use Amazon Web Services for Containerization},
	language = {en-US},
	urldate = {2023-03-05},
	journal = {Amazon Web Services, Inc.},
	file = {Snapshot:/Users/zihengzhang/Zotero/storage/HPY57D32/containerization.html:text/html},
}


@ARTICLE{9758043,
  author={Usman, Muhammad and Badampudi, Deepika and Smith, Chris and Nayak, Himansu},
  journal={IEEE Software}, 
  title={An Ecosystem for the Large-Scale Reuse of Microservices in a Cloud-Native Context}, 
  year={2022},
  volume={39},
  number={5},
  pages={68-75},
  doi={10.1109/MS.2022.3167447}}
  
@book{dobies2020kubernetes,
  title={Kubernetes operators: Automating the container orchestration platform},
  author={Dobies, Jason and Wood, Joshua},
  year={2020},
  publisher={O'Reilly Media}
}

@misc{flux,
	title = {Flux Documentation},
	shorttitle = {Flux Documentation},
	url = {https://fluxcd.io/flux/},
	abstract = {Flux is a tool for keeping Kubernetes clusters in sync with sources of configuration (like Git repositories), and automating updates to configuration when there is new code to deploy.},
	language = {en-US},
	urldate = {2023-04-11},
	journal = {Flux},
	file = {Snapshot:/Users/zihengzhang/Zotero/storage/HPY57D32/containerization.html:text/html},
}
```