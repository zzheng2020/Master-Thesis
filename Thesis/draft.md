# Achieving High Availability in Lifecycle Management of Cloud-Native Network Functions: A Prototype of Database Version Changes

* Content
  1. Intoduction
     1. Overview
     2. Problem Statement
     3. Purpose
     4. Goals
     5. Research Methodology
     6. Delimitations
     7. Sustainability and Ethics
     8. Structure of the thesis
  2. Background
  3. Methodology
  4. Implementation
  5. Evaluation
  6. Conclusion

# 1. Introduction

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

## 2.4 DevOps and Continuous Delivery

DevOps and continuous delivery are interrelated practices that work in tandem to foster a culture of collaboration and automation to optimize software delivery. The goal of these practices is to enhance software delivery by facilitating faster, more reliable, and high-quality releases through automation and collaborative efforts.

DevOps, a portmanteau of "development" and "operations," is an innovative approach that aims to integrate software development and IT operations seamlessly. This interdisciplinary methodology is designed to bridge the gap between these traditionally siloed domains, fostering collaboration, communication, and efficiency. By leveraging a suite of tools, practices, and cultural shifts, DevOps aims to enhance the software development lifecycle (SDLC) through continuous delivery, automated testing, and rapid iteration. As a result, organizations employing DevOps principles can improve the quality, reliability, and speed of their software deployments. Moreover, this holistic approach promotes a culture of shared responsibility, enabling teams to respond more effectively to changing requirements and mitigating risks associated with traditional development methodologies. In essence, DevOps represents a paradigm shift in software engineering, with the potential to reshape organizational structures and drive significant advancements in the field.

Continuous Delivery (CD) is a software engineering approach that emphasizes the automation and streamlining of the software release process, thereby ensuring rapid and reliable deployment of new features and bug fixes. As an integral component of the DevOps methodology, CD fosters a culture of iterative improvement, facilitating seamless integration of code changes into a shared repository, and subsequently delivering these changes to production environments with minimal human intervention.

CD relies on a robust suite of tools and practices, including version control, automated build processes, testing frameworks, and deployment mechanisms, which collectively contribute to reduced lead times and increased deployment frequency. By promoting a systematic, reproducible, and auditable release process, CD mitigates risks associated with manual deployments and fosters greater collaboration among development and operations teams. Furthermore, this approach enables organizations to respond swiftly to changing market demands, enhance customer satisfaction, and maintain a competitive edge. Moreover, Continuous Delivery represents a transformative paradigm within the software engineering landscape, driving innovation and efficiency in an increasingly complex and rapidly evolving technological landscape.

## 2.5 Kubernetes

As previously discussed, Docker proves to be a powerful tool for the development of microservices systems. However, in the context of clusters comprising tens of thousands of containers, numerous challenges arise in terms of scheduling, managing, and dispatching these resources. Consequently, an efficient and flexible management system is required to address these concerns. Kubernetes was developed as a solution to these issues, with its name derived from the Greek word for "helmsman" or "pilot." The platform builds on Google's extensive experience in managing large-scale production workloads and incorporates the best practices and ideas from the community.

Kubernetes is a portable, scalable, open-source platform designed for the management of containerized workloads and services. It enables declarative configuration and automation, while boasting a large, rapidly growing ecosystem of services, support, and tools. Utilizing Kubernetes allows for rapid application deployment, efficient scaling, seamless integration of new application features, resource conservation, and optimization of hardware resource utilization.

Furthermore, Kubernetes offers several key features, including:

1. Portability: Support for public, private, hybrid, and multi-cloud environments.
2. Extensibility: A modular architecture that allows for plug-ins, mounting, and combination of components.
3. Automation: Facilitation of automatic deployment, restart, replication, and scaling/extension.

### 2.5.1 Kubernetes Architecture

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

### 2.5.2 Kubernetes Core Components

Kubernetes encompasses a variety of fundamental components, including Pods, Kubelet, ReplicaSet, and Deployment, which together form the basis for efficient container management and orchestration within the platform.

Kubernetes serves as a container management system, though it does not directly interact with containers. Instead, it operates on the smallest unit called a Pod, which indirectly manages containers. A Master node corresponds to a group of Node nodes, with the Master node responsible for scheduling, network management, controller, and resource object storage. Containers are housed within the Node node, which in turn, stores them inside the Pod. Pods can encapsulate single or multiple containers.

Kubelet is tasked with local Pod maintenance, while kube-proxy manages load balancing between multiple Pods. A Pod is a container that encapsulates other containers created by Docker, functioning as a virtualized group. Pods act as standalone hosts and can house one or more containers. Each Pod possesses its own IP address and hostname, providing an independent sandbox environment.

Pods are typically employed for managing groups of related services during service deployment. They can deploy a single service or a group of related services, which are sets of services connected through a call path of a chain call. Kubernetes excels at managing web service clustering through the control of the number of Pods, while the underlying network of Pods and data storage is facilitated by the creation of Pause containers prior to the establishment of internal containers. Service containers can access localhost, offering high-performance access to local services.

ReplicaSet controllers maintain the desired number of Pod replicas or "service clusters," ensuring consistent replica numbers. In case a Pod service fails, the ReplicaSet controller immediately recreates a new Pod. Replica Controllers utilize tag selectors to maintain a group of related services. In newer versions of Kubernetes, ReplicaSet is recommended over the deprecated ReplicationController.

Deployment objects represent the service deployment architecture model, enabling rolling updates. While ReplicaSet controllers control the number of Pod replicas, they do not support rolling updates. Deployment objects, on the other hand, support rolling updates and are often used in conjunction with ReplicaSets. Deployment manages ReplicaSets, with the latter reestablishing new ReplicaSets and creating new Pods to facilitate project version updates when needed.

### 2.5.3 Stateful vs. Stateless

Stateful and Stateless resources are critical concepts in Kubernetes, particularly when addressing containerized deployment challenges. When deploying PostgreSQL within containers, data loss becomes a significant concern due to the lifecycle of both containers and Pods. Kubernetes does not utilize Deployment for managing stateful services; instead, it employs Deployment for stateless services and StatefulSet for stateful services deployment.

Stateful services are characterized by real-time data storage requirements. In a stateful service cluster, when a service is temporarily removed and later rejoined to the network, the unavailability of the cluster network may present issues. On the other hand, stateless services do not necessitate real-time data storage. Consequently, in a stateless service cluster, temporary removal and reintegration of a service into the network will not impact the clustered service.

StatefulSet is employed to address the challenges associated with using containerized deployment for stateful services. This deployment model is specifically designed for stateful services, ensuring that the Hostname remains unaltered after Pod recreation. This stability enables the Pod to associate data through the Hostname, ultimately improving the management and resilience of stateful services within Kubernetes.

### 2.5.4 Kubernetes Operator

The Kubernetes Operator is a concept designed to encapsulate the knowledge and objectives of an operations professional responsible for managing a service or a set of services. These professionals possess an in-depth understanding of system operations, deployment, and problem resolution. Automation is a preferred method for handling repetitive tasks in the context of Kubernetes workloads, and the Operator pattern serves to embody the task automation code written outside the functionality provided by Kubernetes itself \cite{dobies2020kubernetes}.

Operators function by extending the Kubernetes control plane and API. They introduce an endpoint (referred to as a custom resource or CR) to the Kubernetes API, which also incorporates a control plane component that monitors and maintains new resource types. An Operator comprises a set of Controllers, each listening to specific Kubernetes resources. These Controllers can implement a reconciliation loop, with each Controller being responsible for monitoring a designated resource and triggering reconciliation upon creation, update, or deletion of the monitored resource.

Several open-source projects facilitate the creation of Kubernetes Operators, such as Operator SDK, Kubebuilder, KUDO, and Metacontroller. These tools streamline the development and management of Operators within the Kubernetes ecosystem.

Kubernetes Controllers, on the other hand, are non-terminating loops that regulate the state of the system, striving to align the current state with the desired state as closely as possible through a process known as the Reconciliation loop. In the context of Kubernetes, a set of built-in Controllers operates within the controller-manager located in the master node. These Controllers work in tandem with Operators to ensure efficient and accurate management of resources in a Kubernetes environment.

## 2.6 Flux

Flux is a specialized tool designed to maintain synchronization between Kubernetes clusters and configuration sources, such as Git repositories. This synchronization enables the automation of updates to the configuration when new code deployments arise. Developed with a focus on leveraging Kubernetes' API extension system, Flux seamlessly integrates with core components of the Kubernetes ecosystem, including Prometheus.

One of the key strengths of Flux is its support for multi-tenancy, allowing for the efficient management of multiple tenants within a single cluster. Additionally, Flux facilitates the synchronization of an arbitrary number of Git repositories, providing flexibility and adaptability in managing diverse sources of configuration. This capability ensures that the Kubernetes clusters remain up-to-date with the latest configurations, enhancing the efficiency and reliability of container management and orchestration.

# 3. System Design



# 4. Work

# 5. Results

# 6. Conculsions



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