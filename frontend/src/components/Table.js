import React from 'react';

const Table = ({ transactions, setModalOpen }) => {

  return (
    <table className="table">
      <thead>
        <tr>
          <th>Timestamp</th>
          <th>ID</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        {transactions.map( (transaction, index) => {
          return (
            <tr key={ index } onClick={ ()=>setModalOpen(transaction) }>
              <td>{ (new Date(transaction.Timestamp * 1000)).toLocaleString("en-US") }</td>
              <td>{ transaction.ID }</td>
              <td>{ transaction.Value }</td>
            </tr>
          )
        })}
      </tbody>
    </table>
  );
}

export default Table