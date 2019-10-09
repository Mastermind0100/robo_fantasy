package com.herokuapp.ringo.robofantasy;

public class Matches {
    private int id;
    private String red_team;
    private String blue_team;
    private int status;
    private int category;

    public Matches(int id, String red_team, String blue_team, int status, int category){
        this.id = id;
        this.red_team = red_team;
        this.blue_team = blue_team;
        this.status = status;
        this.category = category;
    }

    public int getId() {
        return id;
    }

    public String getRed_team() {
        return red_team;
    }

    public int getCategory() {
        return category;
    }

    public int getStatus() {
        return status;
    }

    public String getBlue_team() {
        return blue_team;
    }

    public void setId(int id) {
        this.id = id;
    }

    public void setBlue_team(String blue_team) {
        this.blue_team = blue_team;
    }

    public void setCategory(int category) {
        this.category = category;
    }

    public void setRed_team(String red_team) {
        this.red_team = red_team;
    }

    public void setStatus(int status) {
        this.status = status;
    }
}
