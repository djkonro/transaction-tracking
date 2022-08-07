import React from 'react';
import Modal from 'react-modal';


const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      transform: 'translate(-50%, -50%)',
    },
  };
  
  Modal.setAppElement('body');

  const Details = ({ modalIsOpen, setModalClose, modalTransaction }) =>{
    let subtitle;
  
    function afterOpenModal() {
      // references are now sync'd and can be accessed.
      subtitle.style.color = '#000';
    }
  
    function closeModal() {
        setModalClose()
    }
  
    return (
      <div id='modal'>
        <Modal
          isOpen={modalIsOpen}
          onAfterOpen={afterOpenModal}
          onRequestClose={closeModal}
          style={customStyles}
          contentLabel="Transaction Details"
        >
          <h4 ref={(_subtitle) => (subtitle = _subtitle)}>Transaction Details</h4>
          <button onClick={closeModal} type="button" className="close" aria-label="Close">
              <span aria-hidden="true">&times;</span>
          </button>

          <ul>
            <li><b>Transaction ID:</b> { modalTransaction.ID }</li>
            <li><b>Value:</b> { modalTransaction.Value }</li>
            <li><b>Timestamp:</b> { (new Date(modalTransaction.Timestamp * 1000)).toLocaleString("en-US") }</li>
            <li><b>Receiver:</b> { modalTransaction.Receiver }</li>
            <li><b>Sender:</b> { modalTransaction.Sender }</li>
            <li><b>Confirmed:</b> { modalTransaction.Confirmed? "True" : "False" }</li>
          </ul>
          
        </Modal>
      </div>
    );
  }

export default Details;
