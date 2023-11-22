// ResumeBuilder/internal/models/models.go

/*
This file contains the structures for the resume
such as Education, Skills Experience, Projects
and finally a resume structure that utilizes them all
*/

package models
// Resume structure
type Education struct {
	Degree    string
	School    string
	Location  string
	Gpa       string
	StartDate string
	EndDate   string
}

// Experience structure
type Experience struct {
	Company   string
	Title     string
	Location  string
	StartDate string
	EndDate   string
	Bullet1   string
	Bullet2   string
	Bullet3   string
}

// Project structure
type Project struct {
	ProjectName string
	Description string
	ProjectDate string
	ProjectLink string
}
type Resume struct {
	Name       string
	CityState  string
	Email      string
	Website    string
	Education  Education
	Skills     []string
	Experience []Experience
	Projects   []Project
}

// Below are helper functions to generate the structures


func newEducation(d string, s string, g string, sd, ed, lo string ) *Education {
	return &Education{
		Degree:    d,
		School:    s,
		Location:  lo,
		Gpa:       g,
		StartDate: sd,
		EndDate:   ed,
	}
}

func newExperience(c string, t string, sd, ed, b1, b2, b3 string) *Experience {
	return &Experience{
		Company:       c,
		Title:     t,
		StartDate: sd,
		EndDate:   ed,
		Bullet1:   b1,
		Bullet2:   b2,
		Bullet3:   b3,
	}
}

func newProject(n string, d string, da string, l string) *Project {
	return &Project{
		ProjectName: n,
		Description: d,
		ProjectDate: da,
		ProjectLink: l,
	}
}

func newResume(n string, e string, w string, edu Education, sk []string, exp []Experience, proj []Project) *Resume {
	return &Resume{
		Name:	n,
		Email: e,
		Website: w,
		Education: edu,
		Skills: sk,
		Experience: exp,
		Projects: proj,
	}
}