// ResumeBuiler/ResumeBuilder.go

// ResumeBuilder.go is the entry point for the resume builder

package main

import (
	// import the util package from ResumeBuilder/util
	"ResumeBuilder/util"
)

func main() {
	// Generate the resume LaTex file using the utility packages
	r := util.CreateResume()
	util.GenerateResume(r)

}