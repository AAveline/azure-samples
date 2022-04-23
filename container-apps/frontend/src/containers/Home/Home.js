import { useState } from 'react';
import { gql, useQuery, useMutation } from '@apollo/client';
import { Table, Button, Container, Row, Modal, Form } from "react-bootstrap";

import Loader from "../../components/Loader";

const GET_PRODUCTS = gql`
  query {
    listProducts {
      name
      id
    }
  }
`;

const CREATE_PRODUCT = gql`
  mutation($input: CreateProductInput!) {
    createProduct(input: $input)
  }
`;

const DELETE_PRODUCT = gql`
  mutation($input: String!) {
    deleteProduct(input: $input)
  }
`;

function Home() {
  const [product, setProduct] = useState("");  
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const { loading, data } = useQuery(GET_PRODUCTS); 
  
  const [createProduct, { loading: createProductLoading }] = useMutation(CREATE_PRODUCT, {
    refetchQueries: [GET_PRODUCTS],
    onCompleted: () => handleClose()
  });
  const [deleteProduct] = useMutation(DELETE_PRODUCT, {
    refetchQueries: [GET_PRODUCTS],
    onCompleted: () => handleClose()
  });

  return (
    <div className="App">
      <Container>
        <Row>
          <Button variant="primary" size="lg" onClick={handleShow}>Add product</Button>
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
                    data?.listProducts.map((product, index) => (
                      <tr key={index}>
                        <td>{product.name}</td>
                        <td>
                          <Button variant="danger" size="lg" onClick={() => deleteProduct({
                            variables: {
                              input: product.id
                            }
                          })}>Delete</Button>
                        </td>
                      </tr>
                    ))
                  }
                </tbody>
              </Table>
            </Row>
          )
        }
        <Modal show={show} onHide={handleClose}>
          <Modal.Header closeButton>
            <Modal.Title>Product creation</Modal.Title>
          </Modal.Header>
          <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="formBasicName">
              <Form.Label>Product Name</Form.Label>
              <Form.Control type="input" placeholder="Enter a name" onChange={e => setProduct(e.target.value)}/>
            </Form.Group>
          </Form>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleClose}>
              Close
            </Button>
            <Button variant="primary" onClick={() => createProduct({
              variables: {
                input: {
                  name: product
                },
              }
            })} disabled={createProductLoading}>
              Create
            </Button>
          </Modal.Footer>
        </Modal>
       </Container>
    </div>
  );
}

export default Home;
