package com.herokuapp.ringo.robofantasy;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import com.herokuapp.ringo.robofantasy.ui.ui.leaderboard.LeaderboardFragment;

import java.util.List;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

public class ParticipantsAdapter extends ArrayAdapter<Participants> {

    public ParticipantsAdapter(@NonNull Context context, int resource, @NonNull List<Participants> objects) {
        super(context, resource, objects);
    }

    @NonNull
    @Override
    public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
        View listItemView = convertView;
        if(listItemView == null) {
            listItemView = LayoutInflater.from(getContext()).inflate(
                    R.layout.item_message, parent, false);
        }

        TextView mpoints = listItemView.findViewById(R.id.pointss);
        TextView mname = listItemView.findViewById(R.id.name);
        TextView mposition = listItemView.findViewById(R.id.position);

        Participants participants = getItem(position);
        mpoints.setText(Integer.toString(participants.getPoints()));
        mname.setText(participants.getName());
        mposition.setText(Integer.toString(participants.getPosition()));

        return listItemView;
    }
}
