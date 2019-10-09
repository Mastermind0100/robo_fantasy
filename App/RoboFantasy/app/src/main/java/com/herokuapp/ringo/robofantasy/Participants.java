package com.herokuapp.ringo.robofantasy;

public class Participants {
    private int position;
    private String name;
    private int points;

    public Participants(int position, String name, int points){
        this.name = name;
        this.points = points;
        this.position = position;
    }
    public int getPosition(){
        return this.position;
    }

    public int getPoints() {
        return points;
    }

    public String getName() {
        return name;
    }

    public void setPosition(int position) {
        this.position = position;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setPoints(int points) {
        this.points = points;
    }
}
