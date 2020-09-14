import React, {Component} from 'react';
import Grid from "semantic-ui-react/dist/commonjs/collections/Grid";
import {Button, GridColumn, GridRow} from "semantic-ui-react";
import SkillComponent from "../MinorComponents/skillComponent";
import {getSkillsByFieldId} from "../../../Js Functionals/SkillsPage/skillsGetterByFields";

class PersonSkillsComponent extends Component {

    constructor(props, context) {
        super(props, context);

        if (props.fieldId) getSkillsByFieldId(this.setSkills, props.fieldId)
        if (this.props.skillCardObject) this.setState({
            optionsObject: this.props.skillCardObject,
            optionsCard: this.props.skillCardObject.map((object) => this.createCard(object))
        })
    }

    state = {
        optionsCard: [],
        optionsObject: [],
    };

    setSkills = (skills) => {
        this.setState((prevState) => {
            return {
                optionsObject: skills,
                optionsCard: skills.map((value) => this.createCard(value))
            }
        })
    }

    render() {
        return (
            <div>
                {this.createButton()}
                {this.createGrid()}
            </div>

        );
    }

    createButton = () => {
        if (this.props.hasExit) {
            return <Button onClick={() => this.props.exitFunction()} icon='angle left'/>
        }
    }

    createGrid = () => {
        let skillCardArray = this.state.optionsObject;
        let numberOfRows = (skillCardArray.length) / this.props.columnSize;
        if (skillCardArray.length % this.props.columnSize !== 0) {
            numberOfRows += 1;
        }
        let listOfRows = [];
        let counter = 0
        for (let i = 0; i < numberOfRows - 1; i++) {
            let listOfColumns = []
            for (let j = 0; j < this.props.columnSize; j++) {
                let card = this.createCard(skillCardArray[counter])
                listOfColumns[listOfColumns.length] = <GridColumn>{card}</GridColumn>;
                counter += 1;
            }
            listOfRows[listOfRows.length] = <Grid.Row>{listOfColumns}</Grid.Row>;
        }

        let listOfColumns = []
        for (let i = counter; i < skillCardArray.length; i++) {
            listOfColumns[listOfColumns.length] = <GridColumn>
                {skillCardArray[i]}
            </GridColumn>;
        }
        listOfRows[listOfRows.length] = <GridRow>{listOfColumns}</GridRow>;
        return <Grid columns={this.props.columnSize} divided>
            {listOfRows}
        </Grid>;
    }

    createCard = (skill) => {
        return <SkillComponent name={skill.name} type={skill.type}
                               skillDeleter={this.props.skillDeleter} skillAdder={this.props.skillAdder}
                               personHasSkill={this.props.skillIncludes ? this.props.skillIncludes(skill) : true}/>;
    }

}

export default PersonSkillsComponent;
