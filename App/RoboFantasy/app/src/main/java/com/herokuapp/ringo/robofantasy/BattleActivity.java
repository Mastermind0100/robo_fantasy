package com.herokuapp.ringo.robofantasy;

import android.os.Bundle;
import android.view.MenuItem;

import com.google.android.material.bottomnavigation.BottomNavigationView;
import com.herokuapp.ringo.robofantasy.ui.dashboard.DashboardFragment;
import com.herokuapp.ringo.robofantasy.ui.home.HomeFragment;
import com.herokuapp.ringo.robofantasy.ui.notifications.NotificationsFragment;
import com.herokuapp.ringo.robofantasy.ui.ui.leaderboard.LeaderboardFragment;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentTransaction;
import androidx.navigation.NavController;
import androidx.navigation.Navigation;
import androidx.navigation.ui.AppBarConfiguration;
import androidx.navigation.ui.NavigationUI;

public class BattleActivity extends AppCompatActivity {

    private HomeFragment homeFragment;
    private DashboardFragment dashboardFragment;
    UserSessionManager session;
    private LeaderboardFragment leaderboardFragment;
    private NotificationsFragment notificationsFragment;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_battle);
        //User session management



        BottomNavigationView navView = findViewById(R.id.nav_view);
        // Passing each menu ID as a set of Ids because each
        // menu should be considered as top level destinations.
        AppBarConfiguration appBarConfiguration = new AppBarConfiguration.Builder(
                R.id.navigation_home, R.id.navigation_dashboard, R.id.navigation_leaderboard, R.id.navigation_notifications)
                .build();
        NavController navController = Navigation.findNavController(this, R.id.nav_host_fragment);
        NavigationUI.setupActionBarWithNavController(this, navController, appBarConfiguration);
        NavigationUI.setupWithNavController(navView, navController);

        if (savedInstanceState == null) {
            getSupportFragmentManager().beginTransaction().replace(R.id.nav_host_fragment, new HomeFragment()).commit();
        }
        homeFragment = new HomeFragment();
        dashboardFragment = new DashboardFragment();
        leaderboardFragment = new LeaderboardFragment();
        notificationsFragment = new NotificationsFragment();


        navView.setOnNavigationItemSelectedListener(new BottomNavigationView.OnNavigationItemSelectedListener() {
            @Override
            public boolean onNavigationItemSelected(@NonNull MenuItem menuItem) {

                switch (menuItem.getItemId()){
                    case R.id.navigation_home: {
                        setFragment(homeFragment);
                        return true;
                    }
                    case R.id.navigation_dashboard: {
                        setFragment(dashboardFragment);
                        return true;
                    }
                    case R.id.navigation_leaderboard :{
                        setFragment(leaderboardFragment);
                        return true;
                    }
                    case R.id.navigation_notifications :{
                        setFragment(notificationsFragment);
                        return true;
                    }
                    default: return false;
                }

            }
        });
    }

    private void setFragment(Fragment fragment){
        FragmentTransaction fragmentTransaction = getSupportFragmentManager().beginTransaction();
        fragmentTransaction.replace(R.id.nav_host_fragment, fragment);
        fragmentTransaction.addToBackStack(null);
        fragmentTransaction.commit();
    }
}
