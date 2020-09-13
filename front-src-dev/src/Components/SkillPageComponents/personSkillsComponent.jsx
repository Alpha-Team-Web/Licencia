import React, {Component} from 'react';
import Grid from "semantic-ui-react/dist/commonjs/collections/Grid";
import {GridColumn, GridRow} from "semantic-ui-react";

class PersonSkillsComponent extends Component {


    constructor(props, context, state) {
        super(props, context);
        this.state = state
    }

    state;
    /*state={
        skillCardArray
    }*/
    /*props={
        columnSize,

    }*/

    render() {
        return (
            this.createGrid()
        );
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
                ReactDOM.render(skillCardArray[counter], gridColumn);
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
}

export default PersonSkillsComponent;
