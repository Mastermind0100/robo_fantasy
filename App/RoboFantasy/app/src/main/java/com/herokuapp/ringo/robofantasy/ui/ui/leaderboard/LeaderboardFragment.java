package com.herokuapp.ringo.robofantasy.ui.ui.leaderboard;

import androidx.lifecycle.Observer;
import androidx.lifecycle.ViewModelProviders;
import com.herokuapp.ringo.robofantasy.*;
import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import android.provider.Telephony;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ListView;
import android.widget.TextView;

import com.herokuapp.ringo.robofantasy.R;
//
// import com.herokuapp.ringo.robofantasy.ui.R;
import com.herokuapp.ringo.robofantasy.ui.dashboard.DashboardViewModel;

import java.util.ArrayList;

public class LeaderboardFragment extends Fragment {

    private LeaderboardViewModel leaderboardViewModel;

    public View onCreateView(@NonNull LayoutInflater inflater,
                             ViewGroup container, Bundle savedInstanceState) {
        leaderboardViewModel = ViewModelProviders.of(this).get(LeaderboardViewModel.class);
        View root = inflater.inflate(R.layout.leaderboard_fragment, container, false);

        ArrayList<Participants> participants = new ArrayList<Participants>();
        participants.add(new Participants(1,"A",64));
        participants.add(new Participants(2,"b",65));
        participants.add(new Participants(3,"c",655));



        //try{

        final ParticipantsAdapter mparticipantsAdapter = new ParticipantsAdapter(this.getActivity(), R.layout.item_message, participants);
        //QueryUtils queryUtils = new QueryUtils();
        // Find a reference to the {@link ListView} in the layout
        ListView participantListView =  root.findViewById(R.id.list);
        if (participantListView!=null){
        // Create a new {@link ArrayAdapter} of earthquakes
        // Set the adapter on the {@link ListView}
        // so the list can be populated in the user interface
        participantListView.setAdapter(mparticipantsAdapter);}
        else{
            System.out.println("NUll hai");
        }
        /*final TextView textView = root.findViewById(R.id.text_dashboard);
        leaderboardViewModel.getText().observe(this, new Observer<String>() {
            @Override
            public void onChanged(@Nullable String s) {

            }
        });*/

        /*catch (Exception e){
            e.printStackTrace();
        }
*/
        return root;
    }

}
