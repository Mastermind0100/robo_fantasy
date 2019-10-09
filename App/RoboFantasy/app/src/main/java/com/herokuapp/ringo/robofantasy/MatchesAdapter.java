package com.herokuapp.ringo.robofantasy;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import java.util.List;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;

public class MatchesAdapter extends ArrayAdapter<Matches> {

    public MatchesAdapter(@NonNull Context context, int resource, List<Matches> objects) {
        super(context, resource, objects);
    }

    public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
        View listItemView = convertView;
        if (listItemView == null) {
            listItemView = LayoutInflater.from(getContext()).inflate(
                    R.layout.item_matches, parent, false);
        }

        TextView itemno = listItemView.findViewById(R.id.number);
        TextView match = listItemView.findViewById(R.id.matches);

        Matches matches = getItem(position);

        itemno.setText(Integer.toString(position+1));
        match.setText(matches.getBlue_team() + " Vs " + matches.getRed_team());

        return listItemView;
    }
}
