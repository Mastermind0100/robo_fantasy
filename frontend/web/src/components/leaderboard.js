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
                    <tr>
                    <td>2</td>
                    <td>{this.state.data[1].user.username}</td>
                    <td>{this.state.data[1].points}</td>
                    </tr>
                    <tr>
                    <td>3</td>
                    <td>{this.state.data[2].user.username}</td>
                    <td>{this.state.data[2].points}</td>
                    </tr>
                    <tr>
                    <td>4</td>
                    <td>{this.state.data[3].user.username}</td>
                    <td>{this.state.data[3].points}</td>
                    </tr>
                    <tr>
                    <td>5</td>
                    <td>{this.state.data[4].user.username}</td>
                    <td>{this.state.data[4].points}</td>
                    </tr>
                    <tr>
                    <td>6</td>
                    <td>{this.state.data[5].user.username}</td>
                    <td>{this.state.data[5].points}</td>
                    </tr>
                    <tr>
                    <td>7</td>
                    <td>{this.state.data[6].user.username}</td>
                    <td>{this.state.data[6].points}</td>
                    </tr>
                    <tr>
                    <td>8</td>
                    <td>{this.state.data[7].user.username}</td>
                    <td>{this.state.data[7].points}</td>
                    </tr>
                    <tr>
                    <td>9</td>
                    <td>{this.state.data[8].user.username}</td>
                    <td>{this.state.data[8].points}</td>
                    </tr>
                    <tr>
                    <td>10</td>
                    <td>{this.state.data[9].user.username}</td>
                    <td>{this.state.data[9].points}</td>
                    </tr>
                </tbody>
            </Table>
        )
        }
    }
}

export default leaderboard;