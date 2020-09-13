import React, {Component} from 'react';
import PersonSkillsComponent from "./personSkillsComponent";
import ProfileForm from "../ProfilePageComponents/ProfileSectionComponents/ProfileForm";
import {Input} from "semantic-ui-react";
import FieldsListComponent from "./fieldsListComponent";

class SkillsPageComponent extends Component {
    mainDivStyle = {};

    firstRealStyle = {};

    secondRealStyle = {};


    state = {
        personAddedSkills: []
    }
    addSkill = (skill) => {
        this.setState((prevState) => {
            let array = prevState.personAddedSkills;
            if (!this.skillIncludes(array, skill)) array.push(skill)
            return {personAddedSkills: array};
        })
    }
    deleteSkill = (skill) => {
        this.setState((prevState) => {
            return {
                personAddedSkills: prevState.personAddedSkills.filter(value => !this.skillEquals(value, skill))
            }
        })
    }
    skillIncludes = (skill) => this.state.personAddedSkills.filter(value => this.skillEquals(value, skill)).length !== 0;
    skillEquals = (skill1, skill2) => skill1.name === skill2.name && skill1.type === skill2.type;

    render() {
        return (
            <div style={this.mainDivStyle}>
                <div style={this.firstRealStyle}>
                    <PersonSkillsComponent skillDeleter={this.deleteSkill}
                                           personAddedSkills={this.state.personAddedSkills}/>
                </div>

                <div style={this.secondRealStyle}>
                    <Input icon='search' className='ui-rtl'/>
                    <FieldsListComponent cards={getSkillFields()} skillAdder={this.addSkill} skillDeleter={this.deleteSkill}
                                         skillIncludes={this.skillIncludes}/>
                </div>
            </div>
        );
    }
}

export default SkillsPageComponent;
