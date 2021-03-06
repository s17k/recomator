/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package automation

import (
	"google.golang.org/api/cloudresourcemanager/v1"
)

// ListProjects lists the projects IDs for projects user has resourcemanager.projects.get permission
func (s *googleService) ListProjects() ([]string, error) {
	projectsService := cloudresourcemanager.NewProjectsService(s.resourceManagerService)
	var projects []string
	err := projectsService.List().Pages(s.ctx, func(r *cloudresourcemanager.ListProjectsResponse) error {
		for _, project := range r.Projects {
			projects = append(projects, project.ProjectId)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return projects, nil
}
