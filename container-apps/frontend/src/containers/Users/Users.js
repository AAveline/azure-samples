import { useState } from 'react';
import { useQuery, useMutation, gql } from "@apollo/client";
import { Outlet } from "react-router-dom";
import { Table, Button, Container, Row, Modal, Form } from "react-bootstrap";

import Loader from "../../components/Loader";

const GET_USERS = gql`
  query {
    listUsers {
      id
      fullName
    }

    listProducts {
      id
      name
    }
  }
`;

const CREATE_USER = gql`
  mutation($input: CreateUserInput!) {
    createUser(input: $input)
  }
`;

const CREATE_ORDER = gql`
  mutation($input: CreateOrderInput) {
    createOrder(input: $input)
  }
`;

const DELETE_USER = gql`
  mutation($input: String!) {
    deleteUser(input: $input)
  }
`;

function useToggle(state) {
  const [isOpen, setOpen] = useState(state);

  const toggleState = () => {
    setOpen(!isOpen);
  }

  return [isOpen, toggleState];
}


function Users() {
  const [showOrderCreation, toggleOrderCreation] = useToggle(false);
  const [showUserCreation, toggleUserCreation] = useToggle(false);
  const [user, setUser] = useState({});
  const [selectedProduct, setSelectedProduct] = useState(null);

  const { loading, data } = useQuery(GET_USERS);

  const [createUser] = useMutation(CREATE_USER, {
    refetchQueries: [GET_USERS],
    onCompleted: () => {
      toggleUserCreation();
      setUser({})
    }
  });

  const [createOrder] = useMutation(CREATE_ORDER, {
    refetchQueries: [GET_USERS],
    onCompleted: () => {
      toggleOrderCreation();
      setSelectedProduct(null);
    }
  });
  
  const [deleteUser] = useMutation(DELETE_USER, {
    refetchQueries: [GET_USERS],
    onCompleted: () => toggleOrderCreation()
  });

  return (
    <div>
        <Container>
        <Row>
          <Button variant="primary" size="lg" onClick={toggleUserCreation}>Add user</Button>
        </Row>
        {
          loading ? (
            <Loader />
          ) : (
            <Row>
              <Table striped bordered hover>
                <thead>
                  <tr>
                    <th>Name</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {
                    data?.listUsers.map((user, index) => (
                      <tr key={index}>
                        <td>{user.fullName}</td>
                        <td>
                          <Button variant="primary" size="lg" onClick={toggleOrderCreation} style={{ marginRight: 8 }}>Create order</Button>
                          <Button variant="danger" size="lg" onClick={() => deleteUser({
                            variables: {
                              input: user.id
                            }
                          })}>Delete</Button>
                          <Modal show={showOrderCreation} >
                            <Modal.Header closeButton>
                              <Modal.Title>Order creation</Modal.Title>
                            </Modal.Header>
                            <Modal.Body>
                            <Form>
                              <Form.Group className="mb-3" controlId="formBasicName">
                                <Form.Select aria-label="Select a product" onChange={e => setSelectedProduct(e.target.value)}>
                                  {
                                    data?.listProducts.map((product) => (
                                      <option key={product.id} value={product.id}>{product.name}</option>
                                    ))
                                  }
                                </Form.Select>
                              </Form.Group>
                            </Form>
                            </Modal.Body>
                            <Modal.Footer>
                              <Button variant="secondary" onClick={toggleOrderCreation}>
                                Close
                              </Button>
                              <Button variant="primary" onClick={() => createOrder({
                                variables: {
                                  input: {
                                    userId: user.id,
                                    productId: selectedProduct
                                  }
                                }
                              })}>
                                Create order
                              </Button>
                            </Modal.Footer>
                          </Modal>
                        </td>
                      </tr>
                    ))
                  }
                </tbody>
              </Table>
            </Row>
          )
        }
        <Modal show={showUserCreation} >
          <Modal.Header closeButton>
            <Modal.Title>User creation</Modal.Title>
          </Modal.Header>
          <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="formBasicName">
              <Form.Label>User Name</Form.Label>
              <Form.Control type="input" placeholder="Enter a name" onChange={e => setUser({
                ...user,
                fullName: e.target.value,
              })}/>
            </Form.Group>
            <Form.Group className="mb-3" controlId="formBasicEmail">
              <Form.Label>User email</Form.Label>
              <Form.Control type="input" placeholder="Enter an email" onChange={e => setUser({
                ...user,
                email: e.target.value,
              })}/>
            </Form.Group>
          </Form>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={toggleOrderCreation}>
              Close
            </Button>
            <Button variant="primary" onClick={() => createUser({
              variables: {
                input: user
              }
            })}>
              Create user
            </Button>
          </Modal.Footer>
        </Modal>
       </Container>
      <Outlet />  
    </div>
  )
}


export default Users;