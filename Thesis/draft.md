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

  In addition, monolithic architecture is the lack of effective resource isolation. Once a monolithic application is deployed, it exists in either a running or stopped state. In this context, resources are not adequately isolated, and any issues that occur can result in all customers using the application experiencing disruptions. This lack of resource isolation increases the risk of system crashes and can lead to negative consequences for end-users.

### 2.1.2 Microservices

The microservices architecture is designed to decouple solutions by breaking down functionality into discrete services \cite{dragoni_microservices_2017}. Rather than using a monolithic structure, a microservices architecture employs several supporting microservices, which are developed, managed, and iterated on independently. This approach enables engineers to create applications that revolve around domain components, making product delivery easier by leveraging cloud architecture and platform-based deployment, management, and service capabilities. By using decentralised components, engieeers can more efficiently manage the various services that comprise the application.

The microservices architecture provides several benefits over traditional monolithic architecture, including:

* Decoupled Services

  One of the primary benefits of microservices architecture is that there is no coupling between services. This allows for better scalability, as the microservice units of a microservice system have a strong ability to expand horizontally.

* Business Splitting

  Microservices are split by business, with solid service boundaries. This makes it easy to rewrite code for a business without having to understand the entire business.

* Independent Deployments

  Each service unit of a microservice is deployed independently, running in a separate process. This means that modifications and deployments of microservices have no impact on other services, improving system stability and reducing the risk of downtime.

## CAP Theorem



## ADP Plateform



# Reference

```latex
@article{deacon2009model,
  title={Model-view-controller (mvc) architecture},
  author={Deacon, John},
  journal={Online][Citado em: 10 de mar{\c{c}}o de 2006.] http://www. jdl. co. uk/briefings/MVC. pdf},
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
```