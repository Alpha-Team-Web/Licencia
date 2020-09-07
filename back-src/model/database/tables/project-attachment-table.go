package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type ProjectAttachmentTable struct {
	conn *pg.DB
}

func NewProjectAttachment(db *pg.DB) *ProjectAttachmentTable {
	return &ProjectAttachmentTable{conn: db}
}

func (table *ProjectAttachmentTable) AddProjectAttachment(attachment existence.ProjectAttachment) error {
	_, err := table.conn.Model(&attachment).Insert()
	return err
}

func (table *ProjectAttachmentTable) AddAttachmentIdToProject(fileId string, projectId string) error {
	project := existence.Project{}
	if err := table.conn.Model(&project).Where("project_id = ?", projectId).Select(); err != nil {
		return err
	}
	project.FileIds = append(project.FileIds, fileId)
	_, err := table.conn.Model(&project).Column("file_ids").Where("project_id = ?", projectId).Update()
	return err
}

func (table *ProjectAttachmentTable) UpdateProjectAttachment(fileId string, attachment existence.ProjectAttachment) error {
	_, err := table.conn.Model(&attachment).Column("name", "data", "size").Where("file_id = ?", fileId).Update()
	return err
}

func (table *ProjectAttachmentTable) GetProjectAttachments(projectId string) ([]existence.ProjectAttachment, error) {
	files := []existence.ProjectAttachment{}
	err := table.conn.Model(&files).Where("project_id = ?", projectId).Select()
	return files, err
}

func (table *ProjectAttachmentTable) GetProjectAttachmentById(fileId string) (existence.ProjectAttachment, error) {
	attachment := existence.ProjectAttachment{}
	err := table.conn.Model(&attachment).Where("file_id = ?", fileId).Select()
	return attachment, err
}

func (table *ProjectAttachmentTable) IsThereFileWithId(fileId string) (bool, error) {
	files := []existence.ProjectAttachment{}
	if err := table.conn.Model(&files).Where("file_id = ?", fileId).Select(); err != nil {
		return false, err
	}
	if len(files) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
