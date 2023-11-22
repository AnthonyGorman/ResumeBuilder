// ResumeBuilder/util/helper.go

// helper.go contains the helper functions for the resume builder

package util

import (
	// import the models package from internal/models
	"ResumeBuilder/internal/models"

	// import ioutils package from rwestlund
	"io/ioutil"

	// import fmt 
	"fmt"

	//import bufio
	"bufio"

	// import os
	"os"
)

// createResume takes in input from the user and creates a Resume structure to use in the generateResume function

func CreateResume() models.Resume {
	
	// Set variables
	scanner := bufio.NewScanner(os.Stdin)
	var R models.Resume
	R.Skills = make([]string, 2)
	R.Experience = make([]models.Experience, 4)
	R.Projects = make([]models.Project, 2)

	fmt.Println(`
Welcome to the Resume Builder
-------------------------------------------------------------------------------------
The program will start by asking you questions, then it will generate a LaTex resume.
-------------------------------------------------------------------------------------
	`)

	// Gather input for name and contact information
	fmt.Println(`
The following questions regard your name and contact information.
-------------------------------------------------------------------------------------`)
	fmt.Print("Enter your name: ")
	if scanner.Scan() {R.Name = scanner.Text()}
	fmt.Print("Enter your city and state: ")
	if scanner.Scan() {R.CityState = scanner.Text()}
	if R.CityState == "" {
		R.CityState = "N/A"
	}
	fmt.Print("Enter your email: ")
	if scanner.Scan() {R.Email = scanner.Text()}
	fmt.Print("Enter your website: ")
	if scanner.Scan() {R.Website = scanner.Text()}
	
	// Gather input for education
	fmt.Println(`
The following questions are regarding your education.
-------------------------------------------------------------------------------------`)
	fmt.Print("Enter your University: ")
	if scanner.Scan() {R.Education.School = scanner.Text()}
	fmt.Print("Enter your Degree (eg. Bachelor of Science in Computer Science): ")
	if scanner.Scan() {R.Education.Degree = scanner.Text()}
	fmt.Print("Enter your GPA: ")
	if scanner.Scan() {R.Education.Gpa = scanner.Text()}
	fmt.Print("Enter your start date: ")
	if scanner.Scan() {R.Education.StartDate = scanner.Text()}
	fmt.Print("Enter your end date: ")
	if scanner.Scan() {R.Education.EndDate = scanner.Text()}
	fmt.Print("Enter your school's location: ")
	if scanner.Scan() {R.Education.Location = scanner.Text()}

	// Gather input for skills
	fmt.Println(`
The following questions are regarding your skills.
-------------------------------------------------------------------------------------`)
	fmt.Print("Enter your programming languages: ")
	if scanner.Scan() {R.Skills[0] = scanner.Text()}
	fmt.Print("Enter your Developer tools: ")
	if scanner.Scan() {R.Skills[1] = scanner.Text()}

	// Gather input for experience
	fmt.Println(`
The following questions are regarding your experience. The format will be Company, 
Title, Start Date, End Date, Bullet 1, Bullet 2, Bullet 3. Your title is the name of
your position. The bullets are the outcomes of your experience. There will be four 
experience entries. Clubs can also be entered as an experience.
-------------------------------------------------------------------------------------`)
	fmt.Print("Enter your first Company: ")
	if scanner.Scan() {R.Experience[0].Company = scanner.Text()}
	fmt.Print("Enter your job title (eg. Software Engineer): ")
	if scanner.Scan() {R.Experience[0].Title = scanner.Text()}
	fmt.Print("Enter your job location: ")
	if scanner.Scan() {R.Experience[0].Location = scanner.Text()}
	fmt.Print("Enter your start date: ")
	if scanner.Scan() {R.Experience[0].StartDate = scanner.Text()}
	fmt.Print("Enter your end date: ")
	if scanner.Scan() {R.Experience[0].EndDate = scanner.Text()}
	fmt.Print("Bullet 1: ")
	if scanner.Scan() {R.Experience[0].Bullet1 = scanner.Text()}
	fmt.Print("Bullet 2: ")
	if scanner.Scan() {R.Experience[0].Bullet2 = scanner.Text()}
	fmt.Print("Bullet 3: ")
	if scanner.Scan() {R.Experience[0].Bullet3 = scanner.Text()}
	fmt.Print("Enter your second Company: ")
	if scanner.Scan() {R.Experience[1].Company = scanner.Text()}
	fmt.Print("Enter your job title: ")
	if scanner.Scan() {R.Experience[1].Title = scanner.Text()}
	fmt.Print("Enter your job location: ")
	if scanner.Scan() {R.Experience[1].Location = scanner.Text()}
	fmt.Print("Enter your job start date: ")
	if scanner.Scan() {R.Experience[1].StartDate = scanner.Text()}
	fmt.Print("Enter your end date: ")
	if scanner.Scan() {R.Experience[1].EndDate = scanner.Text()}
	fmt.Print("Bullet 1: ")
	if scanner.Scan() {R.Experience[1].Bullet1 = scanner.Text()}
	fmt.Print("Bullet 2: ")
	if scanner.Scan() {R.Experience[1].Bullet2 = scanner.Text()}
	fmt.Print("Bullet 3: ")
	if scanner.Scan() {R.Experience[1].Bullet3 = scanner.Text()}
	fmt.Print("Enter your third Company: ")
	if scanner.Scan() {R.Experience[2].Company = scanner.Text()}
	fmt.Print("Enter your job title: ")
	if scanner.Scan() {R.Experience[2].Title = scanner.Text()}
	fmt.Print("Enter your job location: ")
	if scanner.Scan() {R.Experience[2].Location = scanner.Text()}
	fmt.Print("Enter your job start date: ")
	if scanner.Scan() {R.Experience[2].StartDate = scanner.Text()}
	fmt.Print("Enter your end date: ")
	if scanner.Scan() {R.Experience[2].EndDate = scanner.Text()}
	fmt.Print("Bullet 1: ")
	if scanner.Scan() {R.Experience[2].Bullet1 = scanner.Text()}
	fmt.Print("Bullet 2: ")
	if scanner.Scan() {R.Experience[2].Bullet2 = scanner.Text()}
	fmt.Print("Bullet 3: ")
	if scanner.Scan() {R.Experience[2].Bullet3 = scanner.Text()}
	fmt.Print("Enter your fourth experience (Club or company): ")
	if scanner.Scan() {R.Experience[3].Company = scanner.Text()}
	fmt.Print("Enter your title: ")
	if scanner.Scan() {R.Experience[3].Title = scanner.Text()}
	fmt.Print("Enter the location: ")
	if scanner.Scan() {R.Experience[3].Location = scanner.Text()}
	fmt.Print("Enter your start date: ")
	if scanner.Scan() {R.Experience[3].StartDate = scanner.Text()}
	fmt.Print("Enter your end date: ")
	if scanner.Scan() {R.Experience[3].EndDate = scanner.Text()}
	fmt.Print("Bullet 1: ")
	if scanner.Scan() {R.Experience[3].Bullet1 = scanner.Text()}
	fmt.Print("Bullet 2: ")
	if scanner.Scan() {R.Experience[3].Bullet2 = scanner.Text()}
	fmt.Print("Bullet 3: ")
	if scanner.Scan() {R.Experience[3].Bullet3 = scanner.Text()}

	// Gather input for the projects
	fmt.Println(`
The following questions will ask you about your projects. You will enter your project
name, description, date, and link. It is expected that you have two projects.
-------------------------------------------------------------------------------------`)
	fmt.Print("Enter your first project name: ")
	if scanner.Scan() {R.Projects[0].ProjectName = scanner.Text()}
	fmt.Print("Enter your project description: ")
	if scanner.Scan() {R.Projects[0].Description = scanner.Text()}
	fmt.Print("Enter your project date: ")
	if scanner.Scan() {R.Projects[0].ProjectDate = scanner.Text()}
	fmt.Print("Enter your project link: ")
	if scanner.Scan() {R.Projects[0].ProjectLink = scanner.Text()}
	fmt.Print("Enter your second project name: ")
	if scanner.Scan() {R.Projects[1].ProjectName = scanner.Text()}
	fmt.Print("Enter your project description: ")
	if scanner.Scan() {R.Projects[1].Description = scanner.Text()}
	fmt.Print("Enter your project date: ")
	if scanner.Scan() {R.Projects[1].ProjectDate = scanner.Text()}
	fmt.Print("Enter your project link: ")
	if scanner.Scan() {R.Projects[1].ProjectLink = scanner.Text()}
	fmt.Println("Generating your resume into a LaTex file...")
	return R
}

// generateResume generates the resume by taking in a resume and outputting it to a LaTex file.
func GenerateResume(R models.Resume) {
	// output the resume to a LaTex file
	var out = LATEXINTRO + `
	\begin{document}
    \begin{center}
        % NAME
        {\Huge\scshape ` + R.Name + `} \\
	\large ` + R.CityState + `\\
	\footnotesize
	\href{mailto:` + R.Email + `}
	{\raisebox{-0.2\height}\faEnvelope\  \underline{` + R.Email + `
	}} ~ 
        \href{` + R.Website + `}{\raisebox{-0.2\height}\faGithub\ \underline{` + R.Website + `}}
	\end{center}

	\section{Education}
	\resumeSubHeadingListStart

	% MAIN INFORMATION
	\resumeSubheading
	{` + R.Education.School + `}{` + R.Education.Location + `}
		{` + R.Education.Degree + `}{` + R.Education.StartDate + ` - ` + R.Education.EndDate + `}
		\resumeSubHeadingListEnd


	\section{Experience}
		\resumeSubHeadingListStart
		
			% COMPANY 1
			\resumeSubheading
		{` + R.Experience[0].Company + `}{` + R.Experience[0].Location + `}{` + R.Experience[0].Title + `}
		{` + R.Experience[0].StartDate + ` - ` + R.Experience[0].EndDate + `}
		\resumeItemListStart
			\resumeItem{` + R.Experience[0].Bullet1 + `}
			\resumeItem{` + R.Experience[0].Bullet2 + `}
			\resumeItem{` + R.Experience[0].Bullet3 + `}
		\resumeItemListEnd

		
			\resumeSubheading
		{` + R.Experience[1].Company + `}{` + R.Experience[1].Location + `}{` + R.Experience[1].Title + `}
		{` + R.Experience[1].StartDate + ` - ` + R.Experience[1].EndDate + `}
		\resumeItemListStart
			\resumeItem{` + R.Experience[1].Bullet1 + `}
			\resumeItem{` + R.Experience[1].Bullet2 + `}
			\resumeItem{` + R.Experience[1].Bullet3 + `}
		\resumeItemListEnd

		
	\resumeSubheading
	{` + R.Experience[2].Company + `}{`+ R.Experience[2].Location + `}{` + R.Experience[2].Title + `}
	{` + R.Experience[2].StartDate + ` - ` + R.Experience[2].EndDate + `}
	\resumeItemListStart
	\resumeItem{` + R.Experience[2].Bullet1 + `}
	\resumeItem{` + R.Experience[2].Bullet2 + `}
	\resumeItem{` + R.Experience[2].Bullet3 + `}
	\resumeItemListEnd

		
	\resumeSubheading
	{` + R.Experience[3].Company + `}{` + R.Experience[3].Location + `}{` + R.Experience[3].Title + `}
	{` + R.Experience[3].StartDate + ` - ` + R.Experience[3].EndDate + `}
	\resumeItemListStart
	\resumeItem{` + R.Experience[3].Bullet1 + `}
	\resumeItem{` + R.Experience[3].Bullet2 + `}
	\resumeItem{` + R.Experience[3].Bullet3 + `}
	\resumeItemListEnd

	\resumeSubHeadingListEnd
	
	\section{Personal Projects} 
    \resumeSubHeadingListStart

        % Project 1
        \resumeProjectHeading
		{` + R.Projects[0].ProjectName + `-\href{` + R.Projects[0].ProjectLink + `}{\raisebox{-0.2\height}\ \underline
		{` + R.Projects[0].ProjectLink + `}}}{` + R.Projects[0].ProjectDate + `}
		\resumeItemListStart
		\resumeItem{` + R.Projects[0].Description + `}
		\resumeItemListEnd

		% Project 2
		\resumeProjectHeading
		{` + R.Projects[1].ProjectName + `-\href{` + R.Projects[1].ProjectLink + `}{\raisebox{-0.2\height}\ \underline
		{` + R.Projects[1].ProjectLink + `}}}{` + R.Projects[1].ProjectDate + `}
		\resumeItemListStart
		\resumeItem{` + R.Projects[1].Description + `}
		\resumeItemListEnd
	\resumeSubHeadingListEnd
	\section{Technical Skills}

    \vspace{-7pt}
    \begin{itemize}
    [leftmargin=0.15in, label={}]\small{\item{
        % LANGUAGES
        \textbf{Languages}{: ` + R.Skills[0] + `} \\
		\textbf{Developer Tools}{: ` + R.Skills[1] + `}
	}}
	\end{itemize}
	\end{document}`

	// Write tex to disk
	err := ioutil.WriteFile("resume.tex", []byte(out), 0644)
	if err != nil {
	        fmt.Println("write failed ", err)
	}
	fmt.Println("write success")

	
}


	


