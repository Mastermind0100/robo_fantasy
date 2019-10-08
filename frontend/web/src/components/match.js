import React from 'react';
import {Carousel,Row,Col,Media} from 'react-bootstrap';
import './style.css';

class Match extends React.Component {
    constructor(props){
        super(props);
    }
    ControlledCarousel(){
        const [index, setIndex] = React.useState(0);
        const [direction, setDirection] = React.useState(null);
      
        const handleSelect = (selectedIndex, e) => {
          setIndex(selectedIndex);
          setDirection(e.direction);
        };
    }
    render(){
            return (
              <Carousel activeIndex={this.ControlledCarousel.index} direction={this.ControlledCarousel.direction} onSelect={this.ControlledCarousel.handleSelect}>
                <Carousel.Item>
                    <h1>Previous Match</h1>
                    <Row>
                        <Col>
                            <Media>
                                <Media.Body>
                                    <h3>this.props</h3>
                                </Media.Body>
                                <img
                                    width={64}
                                    height={64}
                                    className="mr-3"
                                    src="holder.js/64x64"
                                    alt="Generic placeholder"
                                />
                            </Media>
                        </Col>
                        <Col>
                            <Media>
                                <img
                                    width={64}
                                    height={64}
                                    className="mr-3"
                                    src="holder.js/64x64"
                                    alt="Generic placeholder"
                                />
                            <Media.Body>
                                    <h3>this.props</h3>
                            </Media.Body>
                            </Media>
                        </Col>
                    </Row>
                </Carousel.Item>
                <Carousel.Item>
                <h1>Current Match</h1>
                    <Row>
                        <Col>
                            <Media>
                                <Media.Body>
                                    <h3>this.props</h3>
                                </Media.Body>
                                <img
                                    width={64}
                                    height={64}
                                    className="mr-3"
                                    src="holder.js/64x64"
                                    alt="Generic placeholder"
                                />
                            </Media>
                        </Col>
                        <Col>
                            <Media>
                                <img
                                    width={64}
                                    height={64}
                                    className="mr-3"
                                    src="holder.js/64x64"
                                    alt="Generic placeholder"
                                />
                            <Media.Body>
                                    <h3>this.props</h3>
                            </Media.Body>
                            </Media>
                        </Col>
                    </Row>
                </Carousel.Item>
                <Carousel.Item>
                    <h1>Next Match</h1>
                    <Row>
                        <Col>
                            <Media>
                                <Media.Body>
                                    <h3>this.props</h3>
                                </Media.Body>
                                <img
                                    width={64}
                                    height={64}
                                    className="mr-3"
                                    src="holder.js/64x64"
                                    alt="Generic placeholder"
                                />
                            </Media>
                        </Col>
                        <Col>
                            <Media>
                                <img
                                    width={64}
                                    height={64}
                                    className="mr-3"
                                    src="holder.js/64x64"
                                    alt="Generic placeholder"
                                />
                            <Media.Body>
                                    <h3>this.props</h3>
                            </Media.Body>
                            </Media>
                        </Col>
                    </Row>
                </Carousel.Item>
              </Carousel>
            );
        
    }
}

export default Match;