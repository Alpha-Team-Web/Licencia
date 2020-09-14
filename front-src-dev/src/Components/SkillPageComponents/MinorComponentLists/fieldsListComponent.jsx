import React, {Component, Fragment} from 'react';
import {GridColumn, GridRow, Transition} from "semantic-ui-react";
import PersonSkillsComponent from "./personSkillsComponent";
import FieldCardComponent from "../MinorComponents/fieldCardComponent";
import Grid from "semantic-ui-react/dist/commonjs/collections/Grid";
import ReactDOM from "react-dom";


const defaultColumnsSize = 3;

class FieldsListComponent extends Component {

    constructor(props, context) {
        super(props, context);
    }

    cards = this.props.cards

    cardsToCardComponents = (card, index) => <FieldCardComponent skillName={card.skillName}
                                                                 onClick={() => this.setVisible(index, true)}/>
    cardsToCardContainers = (card) => {
        return {
            card: card,
            visible: false,
        }
    }

    state = {
        options: this.cards.map((card) => this.cardsToCardContainers(card)),
    }
    setVisible = (index, visibility) => {
        this.setState((prevState) => {
            prevState.options[index].visible = visibility;
            return prevState;
        })
    }

    splitCards(cards, splitNum) {
        let splitCards = [];
        for (let i = 0; i < cards.length; i++) {
            let splitCardsUnit = [];
            for (let jSplit = 0, jCards = splitNum * i; jSplit < splitNum && jCards < cards.length; jSplit++, jCards++) {
                splitCardsUnit[jSplit] = cards[jCards];
            }
            splitCards[i] = splitCardsUnit;
        }
        return splitCards;
    }

    createGridColumn = (card, index) => <GridColumn>{this.cardsToCardComponents(card, index)}</GridColumn>
    createGridRow = (cards, index) =>
        <Grid.Row>{cards.map((value, index2) => this.createGridColumn(value, index + index2))}</Grid.Row>
    createGrid = (cards) => {
        let columnSize = this.chooser(this.props.columnSize, defaultColumnsSize)
        return (
            <Grid columns={columnSize} divided>
                {this.splitCards(cards, columnSize).map((value, index) => this.createGridRow(value, index * columnSize))}
            </Grid>
        )
    }
    chooser = (choice, defaultValue) => choice ? choice : defaultValue;

    transitioner = (card, index) => {
        return (
            <Transition visible={card.visible} animation='scale' duration={500}>
                <PersonSkillsComponent fieldId={card.id}
                                       skillIncludes={this.props.skillIncludes} skillAdder={this.props.skillAdder}
                                       skillDeleter={this.props.skillDeleter} hasExit
                                       exitFunction={() => this.setVisible(index, false)}/>
            </Transition>
        )
    }

    render() {
        return (
            <div>
                {this.createGrid(this.cards)}
                {this.state.options.map((value) => this.transitioner(value))}
            </div>
        );
    }

}

export default FieldsListComponent;
