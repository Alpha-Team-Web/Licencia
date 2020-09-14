import React, {Component} from 'react';
import Grid from "semantic-ui-react/dist/commonjs/collections/Grid";
import {Button, GridColumn, GridRow} from "semantic-ui-react";
import ReactDOM from 'react-dom';
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
        let grid = <Grid columns={this.props.columnSize} divided/>
        let skillCardArray = this.state.optionsObject;
        let numberOfRows = (skillCardArray.length) / this.props.columnSize;
        if (skillCardArray.length % this.props.columnSize !== 0) {
            numberOfRows += 1;
        }
        let listOfRows = [];
        let counter = 0
        for (let i = 0; i < numberOfRows - 1; i++) {
            let gridRow = <Grid.Row/>
            let listOfColumns = []
            console.log("2333jfseofiesifjseoifj")
            for (let j = 0; j < this.props.columnSize; j++) {
                let gridColumn = <GridColumn/>
                let card = this.createCard(skillCardArray[counter])
                ReactDOM.render(card, gridColumn);
                counter += 1;
                listOfColumns[listOfColumns.length] = gridColumn;
                console.log("1222jfsoifjesiofjsiofj")
            }
            ReactDOM.render(listOfColumns, gridRow);
            listOfRows[listOfRows.length] = gridRow;
        }
        let gridRow = <GridRow/>
        let listOfColumns = []
        for (let i = counter; i < skillCardArray.length; i++) {
            let gridColumn = <GridColumn/>
            console.log("3444fjosiefjoeisjfiosfjseoifj")
            ReactDOM.render(skillCardArray[i], gridColumn);
            listOfColumns[listOfColumns.length] = gridColumn;
        }
        ReactDOM.render(listOfColumns, gridRow);
        listOfRows[listOfRows.length] = gridRow;
        ReactDOM.render(listOfRows, grid);
        return grid;
    }

    createCard = (skill) => {
        return <SkillComponent name={skill.name} type={skill.type}
                               skillDeleter={this.props.skillDeleter} skillAdder={this.props.skillAdder}
                               personHasSkill={this.props.skillIncludes ? this.props.skillIncludes(skill) : true}/>;
    }

}

export default PersonSkillsComponent;
