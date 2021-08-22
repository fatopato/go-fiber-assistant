import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: []
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    // console.log("pRINTING task", this.state.task);
    if (task) {
      axios
        .post(
          endpoint + "/api/v1/todo",
          {
            task
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
              "Access-Control-Allow-Origin": "*"
            }
          }
        )
        .then(res => {
          res.header("Access-Control-Allow-Origin", "*");
          this.getTask();
          this.setState({
            task: ""
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/v1/todo").then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          items: res.data.map(item => {
            console.log(item)
            let color = "yellow";

            if (item.Completed) {
              color = "green";
            }
            return (
              <Card key={item.ID} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.Title}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="green"
                      onClick={() => this.updateTask(item.ID)}
                    />
                    <span style={{ paddingRight: 10 }}>Done</span>
                    <Icon
                      name="undo"
                      color="yellow"
                      onClick={() => this.undoTask(item.ID)}
                    />
                    <span style={{ paddingRight: 10 }}>Undo</span>
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => this.deleteTask(item.ID)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };

  updateTask = id => {
    axios
      .put(endpoint + "/api/v1/todo/complete/" + id, {
        crossDomain:true,
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "Access-Control-Allow-Origin": "*"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };  

  undoTask = id => {
    axios
      .put(endpoint + "/api/v1/todo/undo/" + id, {
        crossDomain:true,
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "Access-Control-Allow-Origin": "*"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  deleteTask = id => {
    axios
        .delete(endpoint + "/api/v1/todo/" + id, {},  {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "Access-Control-Allow-Origin": "*"
        }
        
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };
  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2">
            TO DO LIST
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
            {/* <Button >Create Task</Button> */}
          </Form>
        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;
