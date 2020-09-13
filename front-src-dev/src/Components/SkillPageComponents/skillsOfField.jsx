import React, {Component} from 'react';
import Grid from "semantic-ui-react/dist/commonjs/collections/Grid";
import {GridColumn, GridRow} from "semantic-ui-react";
import ReactDOM from 'react-dom';

class PersonSkillsComponent extends Component {


    constructor(props, context) {
        super(props, context);
        this.setState({
            optionsObject:this.props.skillCardObject,
            optionsCard:this.props.skillCardObject.map((object)=>this.createCard(object))
        })
    }

    state={
        optionsCard:[],
        optionsObject:[],


    };

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

    createCard=(object)=>{
        let card = <skillComponent onDelete={()=>this.deleteCard(object)} contect={object.contentType} skillType={object.skillType}/>
        return card;
    }

    deleteCard=(object)=>{
        let optionsObjectMock = this.state.optionsObject.filter((obj)=>{
            return (obj.contentText!==object.contentText || obj.skillType!==object.skillType)
        })
        this.setState((prevState)=> {
            return {
                optionsObject:optionsObjectMock,
                optionsCard:optionsObjectMock.map((obj)=>this.createCard(obj))
            }
        })
    }
}

export default PersonSkillsComponent;
