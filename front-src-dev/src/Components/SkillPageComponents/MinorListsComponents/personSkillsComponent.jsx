import React, {Component} from 'react';
import Grid from "semantic-ui-react/dist/commonjs/collections/Grid";
import {Button, GridColumn, GridRow} from "semantic-ui-react";
import ReactDOM from 'react-dom';

class PersonSkillsComponent extends Component {


    constructor(props, context) {
        super(props, context);
        this.setState({
            optionsObject: this.props.skillCardObject,
            optionsCard: this.props.skillCardObject.map((object) => this.createCard(object))
        })
    }

    state = {
        optionsCard: [],
        optionsObject: [],


    };

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
        let skillCardArray = this.state.skillCardArray;
        let numberOfRows = (skillCardArray.length) / this.props.columnSize;
        if (skillCardArray.length % this.props.columnSize !== 0) {
            numberOfRows += 1;
        }
        let listOfRows = [];
        let counter = 0
        for (let i = 0; i < numberOfRows - 1; i++) {
            let gridRow = <Grid.Row/>
            let listOfColumns = []
            for (let j = 0; j < this.props.columnSize; j++) {
                let gridColumn = <GridColumn/>
                let card = this.createCard(skillCardArray[counter])
                ReactDOM.render(card, gridColumn);
                counter += 1;
                listOfColumns[listOfColumns.length] = gridColumn;
            }
            ReactDOM.render(listOfColumns, gridRow);
            listOfRows[listOfRows.length] = gridRow;
        }
        let gridRow = <GridRow/>
        let listOfColumns = []
        for (let i = counter; i < skillCardArray.length; i++) {
            let gridColumn = <GridColumn/>
            ReactDOM.render(skillCardArray[i], gridColumn);
            listOfColumns[listOfColumns.length] = gridColumn;
        }
        ReactDOM.render(listOfColumns, gridRow);
        listOfRows[listOfRows.length] = gridRow;
        ReactDOM.render(listOfRows, grid);
        return grid;
    }

    createCard = (skill) => {
        return <skillComponent name={skill.name} type={skill.type}
                               skillDeleter={this.props.skillDeleter} skillAdder={this.props.skillAdder}
                               personHasSkill={this.props.skillIncludes ? this.props.skillIncludes(skill) : true}/>;
    }

}

export default PersonSkillsComponent;
