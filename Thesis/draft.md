# Achieving High Availability in Lifecycle Management of Cloud-Native Network Functions: A Case Study of Database Version Changes

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

# Introduction

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

## Container

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

Finally, agility is an important aspect of containerization. Containerized applications operate in a separate computing environment, enabling software developers to troubleshoot and make changes to the application code without disrupting the operating system, hardware, or other application services. This accelerates software release cycles and enables quick updates using the container model.

## ADP Plateform

Ericsson has developed the Application Development Platform (ADP) ecosystem to facilitate the development of cloud-native applications and services \cite{9758043}. This platform offers extensive technical, process, and organizational support to developers who create cloud network functions or cloud-native applications within Ericsson. The ecosystem is equipped with various tools, processes, and guidelines that enable developers to build top-quality applications that align with the business's needs.

The ADP ecosystem's architecture is based on modern microservices and container principles, which are widely utilized in contemporary software development. The microservices approach enables developers to deconstruct complex applications into smaller, more manageable components that can be developed and deployed independently. Containers offer a lightweight approach to bundling these components, making them more manageable and easier to move between different environments.

Furthermore, the ADP ecosystem provides critical support for the large-scale reuse of microservices across various applications within Ericsson. This feature helps developers reduce development time and costs by enabling them to create microservices once and then reuse them across multiple applications. The ADP Program also offers numerous resources, such as templates, examples, FAQs, and tutorials, to guide developers in utilizing the platform effectively.

The ADP ecosystem is an excellent platform for developers seeking to create cloud-native applications or cloud network functions within Ericsson. Its comprehensive support for technical, process and organizational aspects of software development makes it an ideal choice for teams looking to streamline their development processes while maintaining high levels of quality and efficiency.

## Continuous Delivery and DevOps



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
```