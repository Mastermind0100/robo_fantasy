import React from 'react';
import {Table, Toast} from 'react-bootstrap';
import './style.css';
import test from './test.js'

class leaderboard extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            data : [],
            isLoaded : false,
        }
    }
        
    componentDidMount() {
        fetch('http://robo.smoketrees.dev/leaderboard').then(res => res.json())
        .then(json=> {
            console.log(json)
            this.setState({
                isLoaded: true,
                data: json.data,
            })
        })
    };

    render() {

        var {isLoaded,data}=this.state;

        if(!isLoaded){
            return <h1>loading....</h1>
        }

        else{
        return(
            <Table striped bordered hover variant="dark" align="center" border="primary" style={{width:'90%'}}>
                <thead>
                    <tr>
                    <th>Rank</th>
                    <th>User Name</th>
                    <th>Points</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                    <td>1</td>
                    <td>{this.state.data[0].user.username}</td>
                    <td>{this.state.data[0].points}</td>
                    </tr>
                </tbody>
            </Table>
        )
        }
    }
}

export default leaderboard;