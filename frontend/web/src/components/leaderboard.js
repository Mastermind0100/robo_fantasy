import React from 'react';
import {Table, Toast} from 'react-bootstrap';
import './style.css';

class leaderboard extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            data: null
        };
    }
    
    componentDidMount() {
        const url="..";
        const response = await fetch(url);
        const data1 = await response.json();
        this.setState({data: data1})
    };
    

    render() {
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
                    <td>data.topten[0].user.firstName</td>
                    <td>data.topten[0].points</td>
                    </tr>
                    <tr>
                    <td>2</td>
                    <td>data.topten[1].user.firstName</td>
                    <td>data.topten[1].points</td>
                    </tr>
                    <tr>
                    <td>3</td>
                    <td>data.topten[2].user.firstName</td>
                    <td>data.topten[2].points</td>
                    </tr>
                    <tr>
                    <td>4</td>
                    <td>data.topten[3].user.firstName</td>
                    <td>data.topten[3].points</td>
                    </tr>
                    <tr>
                    <td>5</td>
                    <td>data.topten[4].user.firstName</td>
                    <td>data.topten[4].points</td>
                    </tr>
                    <tr>
                    <td>6</td>
                    <td>data.topten[5].user.firstName</td>
                    <td>data.topten[5].points</td>
                    </tr>
                    <tr>
                    <td>7</td>
                    <td>data.topten[6].user.firstName</td>
                    <td>data.topten[6].points</td>
                    </tr>
                    <tr>
                    <td>8</td>
                    <td>data.topten[7].user.firstName</td>
                    <td>data.topten[7].points</td>
                    </tr>
                    <tr>
                    <td>9</td>
                    <td>data.topten[8].user.firstName</td>
                    <td>data.topten[8].points</td>
                    </tr>
                    <tr>
                    <td>10</td>
                    <td>data.topten[9].user.firstName</td>
                    <td>data.topten[9].points</td>
                    </tr>
                </tbody>
            </Table>
        )
    }
}

export default leaderboard;