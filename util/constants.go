// ResumeBuilder/util/constants.go

// contants.go contains importaint constants for the resume builder


package util

const (

	LATEXINTRO =
	`%-------------------------
	% Resume in Latex
	% Author : Gabriel Sison
	% Based off of: https://github.com/sb2nov/resume
	% License : MIT
	%------------------------
	
	
	\documentclass[letterpaper,11pt]{article}
	\usepackage{latexsym}
	\usepackage[empty]{fullpage}
	\usepackage{titlesec}
	\usepackage{marvosym}
	\usepackage[usenames,dvipsnames]{color}
	\usepackage{verbatim}
	\usepackage{enumitem}
	\usepackage[hidelinks]{hyperref}
	\usepackage{fancyhdr}
	\usepackage[english]{babel}
	\usepackage{tabularx}
	\usepackage{fontawesome5}
	\usepackage{multicol}
	\setlength{\multicolsep}{-3.0pt}
	\setlength{\columnsep}{-1pt}
	\input{glyphtounicode}
	\pagestyle{fancy}
	\fancyhf{} 
	\fancyfoot{}
	\renewcommand{\headrulewidth}{0pt}
	\renewcommand{\footrulewidth}{0pt}
	% Adjust margins
	\addtolength{\oddsidemargin}{-0.6in}
	\addtolength{\evensidemargin}{-0.5in}
	\addtolength{\textwidth}{1.19in}
	\addtolength{\topmargin}{-.7in}
	\addtolength{\textheight}{1.4in}
	\urlstyle{same}
	\raggedbottom
	\raggedright
	\setlength{\tabcolsep}{0in}
	% Sections formatting
	\titleformat{\section}{
	  \vspace{-7pt}\scshape\raggedright\Large\bfseries
	}{}{0em}{}[\color{black}\titlerule \vspace{0pt}]
	\pdfgentounicode=1
	
	
	\newcommand{\resumeItem}[1]{
	  \item\normalsize{
		{#1 \vspace{-3pt}}
	  }
	}
	\newcommand{\resumeSubheading}[4]{
	  \vspace{-3pt}\item
		\begin{tabular*}{1.0\textwidth}[t]{l@{\extracolsep{\fill}}r}
		  \textbf{#1} & \textbf{\small #2} \\
		  \textit{\small#3} & \textit{\normalsize #4} \\
		\end{tabular*}\vspace{-7pt}
	}
	\newcommand{\resumeSubheadingContinue}[2]{
	  \vspace{-3pt}
		\begin{tabular*}{1.0\textwidth}[t]{l@{\extracolsep{\fill}}r}
		  \textit{\small#1} & \textit{\small #2} \\
		\end{tabular*}\vspace{-7pt}
	}
	\newcommand{\resumeProjectHeading}[2]{
	  \vspace{-3pt}\item
		\begin{tabular*}{1.0\textwidth}[t]{l@{\extracolsep{\fill}}r}
		  \textbf{#1} & \textbf{\small #2} \\
		\end{tabular*}\vspace{-7pt}
	}
	\newcommand{\resumeSubItem}[1]{\resumeItem{#1}\vspace{0pt}}
	\renewcommand\labelitemi{$\vcenter{\hbox{\tiny$\bullet$}}$}
	\renewcommand\labelitemii{$\vcenter{\hbox{\tiny$\bullet$}}$}
	\newcommand{\resumeSubHeadingListStart}{\begin{itemize}[leftmargin=0.0in, label={}]}
	\newcommand{\resumeSubHeadingListEnd}{\end{itemize}}
	\newcommand{\resumeItemListStart}{\begin{itemize}}
	\newcommand{\resumeItemListEnd}{\end{itemize}\vspace{0pt}}
	`
)