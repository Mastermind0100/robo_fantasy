package com.herokuapp.ringo.robofantasy.ui.home;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ListView;
import android.widget.TextView;

import androidx.annotation.Nullable;
import androidx.annotation.NonNull;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.Observer;
import androidx.lifecycle.ViewModelProviders;

import com.herokuapp.ringo.robofantasy.Matches;
import com.herokuapp.ringo.robofantasy.MatchesAdapter;
import com.herokuapp.ringo.robofantasy.Participants;
import com.herokuapp.ringo.robofantasy.ParticipantsAdapter;
import com.herokuapp.ringo.robofantasy.R;

import java.util.ArrayList;

public class HomeFragment extends Fragment {

    private HomeViewModel homeViewModel;

    public View onCreateView(@NonNull LayoutInflater inflater,
                             ViewGroup container, Bundle savedInstanceState) {


        homeViewModel = ViewModelProviders.of(this).get(HomeViewModel.class);
        View root = inflater.inflate(R.layout.fragment_home, container, false);

        ArrayList<Matches> list_matches = new ArrayList<Matches>();
        list_matches.add(new Matches(1,"Raven","Touro",200,60));
        list_matches.add(new Matches(2,"Bose","Tanaji",200,60));
        list_matches.add(new Matches(3,"Raven","Tanaji",200,60));


        final MatchesAdapter matchesAdapter = new MatchesAdapter(this.getActivity(), R.layout.item_matches, list_matches);

        ListView matchesListView =  root.findViewById(R.id.list_matches);
        if (matchesListView!=null){

            matchesListView.setAdapter(matchesAdapter);}
        else{
            System.out.println("NUll hai");
        }

        return root;
    }
}