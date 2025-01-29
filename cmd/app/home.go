package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sullyh7/portfolio/view/components"
	"github.com/sullyh7/portfolio/view/home"
	"github.com/sullyh7/portfolio/view/layout"
)

func (a *application) homeHandler(c echo.Context) error {
	return a.render(c, http.StatusOK, home.Index(
		layout.AppProps{
			NavLinks: []components.NavLink{
				{Title: "Projects", Href: "#projects"},
				{Title: "Experience", Href: "#experience"},
				{Title: "Blog", Href: "/blog"},
			},
		},
		home.HomePageProps{
			AboutMe: components.HeaderProps{
				Name: "Sulayman Hasan",
				Role: "Software Engineer & Computer Science Student",
				Description: `
					Driven and detail oriented Computer Science graduate (2025) with hands on experience with software development, web development and microcontroller programming.
				`,
				Socials: []components.Social{
					{Name: "GitHub", URL: "https://github.com/sulaymanhasan"},
					{Name: "LinkedIn", URL: "https://uk.linkedin.com/in/sulaymanhasan"},
				},
			},
			Experience: []components.ExperienceProps{
				{
					Company:     "EE",
					Role:        "Retail Guide",
					Years:       "10/2023 – Present",
					Description: "Guided customers through the mobile purchasing process and provided tailored recommendations.",
				},
				{
					Company:     "Sports Direct",
					Role:        "Retail Assistant",
					Years:       "10/2021 – 05/2023",
					Description: "Provided excellent customer service and consistently exceeded monthly commission targets.",
				},
			},
			Projects: []components.ProjectCardProps{
				{
					Name:         "Smart Plant Growth Monitoring System",
					Date:         "09/2024 – Present",
					ProjectType:  "University Final Year Project",
					Status:       "In Progress",
					Description:  "Built a microcontroller-based system to monitor soil nutrients, moisture, and light data. Integrated OpenAI API for tailored gardening advice.",
					Tech:         []string{"Arduino", "Golang", "PostgreSQL", "SwiftUI"},
					Contribution: "Developed microcontroller logic, built a Golang server with REST API, and designed an iOS app.",
					URL:          "",
					LiveDemo:     false,
				},
				{
					Name:         "Multilingual Blog for Arabic Learning",
					Date:         "08/2024 – Present",
					ProjectType:  "Personal Project",
					Status:       "In Progress",
					Description:  "A full-stack language learning blog built with Golang, Templ, and HTMX.",
					Tech:         []string{"Golang", "HTMX", "Templ", "OAuth"},
					Contribution: "Designed and developed the entire platform, including authentication and dynamic content rendering.",
					URL:          "https://langblog.online",
					LiveDemo:     true,
				},
				{
					Name:         "Web Scraper for Home Be",
					Date:         "07/2024",
					ProjectType:  "Freelance",
					Status:       "Completed",
					Description:  "Automated stock data retrieval from a supplier's website and pushed results to a WordPress API.",
					Tech:         []string{"Python", "BeautifulSoup", "CSV"},
					Contribution: "Built the scraper and automated the integration with WordPress.",
					URL:          "",
					LiveDemo:     false,
				},
				{
					Name:         "CVHawk Frontend Development",
					Date:         "07/2024",
					ProjectType:  "Freelance",
					Status:       "Completed",
					Description:  "Collaborated with Bluerelic LTD to develop the frontend of CVHawk.",
					Tech:         []string{"HTML", "CSS", "JavaScript"},
					Contribution: "Developed the user interface and optimized responsiveness.",
					URL:          "https://cvhawk.com",
					LiveDemo:     true,
				},
				{
					Name:         "Pet Minder Booking Website",
					Date:         "01/2024 – 04/2024",
					ProjectType:  "University Course Project",
					Status:       "Completed",
					Description:  "Developed a web platform to connect pet owners with pet minders, featuring live chat and secure payment integration.",
					Tech:         []string{"Next.js", "Supabase", "TypeScript"},
					Contribution: "Wrote domain analysis reports, designed UML diagrams, and built the platform with cutting-edge technologies.",
					URL:          "",
					LiveDemo:     false,
				},
			},
		},
	))
}
