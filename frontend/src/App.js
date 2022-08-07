import React, { Component } from 'react';

import Search from './components/Search';
import Table from './components/Table';
import Details from './components/Details';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      transactions: [],
      modalIsOpen: false,
      modalTransaction: {}
    }
  }

  componentDidMount() {
    let transactionPolling = () => {
      fetch('http://localhost:5000/api/transactions')
      .then(res => res.json())
      .then(json => json.transactions)
      .then((transactions) => {
        this.setState({ 'transactions': transactions })
        setTimeout(transactionPolling, 10000);
      })
    }
    setTimeout(transactionPolling, 1000);
  }

  render() {

    let query = new URL(window.location).searchParams.get('search')

    let setModalClose = () => {
      this.setState({ 'modalIsOpen': false })
    }

    let setModalOpen = (transaction) => {
      this.setState({ 'modalIsOpen': true })
      this.setState({ 'modalTransaction': transaction })
    }

    let filterTransactions = (transactions, transaction) => {
      if (!query) {
          return transactions;
      }

      return transactions.filter((transaction) => {
          let transactionData = transaction.ID.toLowerCase()
          transactionData += transaction.Value
          transactionData += transaction.Timestamp
          transactionData += transaction.Receiver.toLowerCase()
          transactionData += transaction.Sender.toLowerCase()

          console.log(transactionData)
          return transactionData.includes(query)
      });
    };

    return (
      <div className="App">
        <Details modalIsOpen={ this.state.modalIsOpen} setModalClose={ setModalClose } modalTransaction={ this.state.modalTransaction } />
         <Search />
        <h4>
          Transactions table
        </h4>
        <Table transactions={ filterTransactions(this.state.transactions, query) } setModalOpen={ setModalOpen }/>
      </div>
    );
  }
}

export default App;
