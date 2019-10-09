package com.herokuapp.ringo.robofantasy.ui.ui.leaderboard;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

public class LeaderboardViewModel extends ViewModel {
    // TODO: Implement the ViewModel
    private MutableLiveData<String> mText;

    public LeaderboardViewModel() {
        mText = new MutableLiveData<>();
        mText.setValue("This is Leaderboard fragment");
    }

    public LiveData<String> getText() {
        return mText;
    }
}
