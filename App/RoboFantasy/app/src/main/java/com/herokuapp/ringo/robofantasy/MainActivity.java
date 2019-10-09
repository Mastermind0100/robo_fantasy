package com.herokuapp.ringo.robofantasy;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.text.Editable;
import android.text.TextUtils;
import android.text.TextWatcher;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;

import com.google.android.material.textfield.TextInputEditText;

import java.util.regex.Matcher;
import java.util.regex.Pattern;


public class MainActivity extends AppCompatActivity {

    private Button btn;
    private Button mFirstTimeUsers;
    private TextInputEditText email;
    UserSessionManager session;
    private TextInputEditText pass;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        session = new UserSessionManager(getApplicationContext());

        email = findViewById(R.id.username);
        pass = findViewById(R.id.password);
        mFirstTimeUsers = findViewById(R.id.btn_ftusers);
        btn = findViewById(R.id.click);

        mFirstTimeUsers.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Intent intent = new Intent(MainActivity.this, RegisterActivity.class);
                startActivity(intent);
                finish();
            }
        });


        btn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                String RegisterID = email.getText().toString();
                String password = pass.getText().toString();

                if (TextUtils.isEmpty(RegisterID)) {
                    Toast.makeText(getApplicationContext(), "Enter email address!", Toast.LENGTH_SHORT).show();
                    return;
                }

                if (TextUtils.isEmpty(password)) {
                    Toast.makeText(getApplicationContext(), "Enter password!", Toast.LENGTH_SHORT).show();
                    return;
                }
                if (!isValidEmailAddress(RegisterID)){
                    Toast.makeText(getApplicationContext(), "Enter Valid Email address!", Toast.LENGTH_SHORT).show();
                    return;
                }
                if(RegisterID.equals("admin@gmail.com") && password.equals("admin")){

                    // Creating user login session
                    // Statically storing name="Android Example"
                    // and email="androidexample84@gmail.com"
                    session.createUserLoginSession("Android Example",
                            "androidexample84@gmail.com");

                    // Starting BattleActivity
                    Intent i = new Intent(getApplicationContext(), BattleActivity.class);
                    i.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP);

                    // Add new Flag to start new Activity
                    i.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
                    startActivity(i);

                    finish();

                }else{

                    // username / password doesn't match&
                    Toast.makeText(getApplicationContext(),
                            "Username/Password is incorrect",
                            Toast.LENGTH_LONG).show();

                }
                /*else
                startActivity(new Intent(MainActivity.this, BattleActivity.class));*/
            }
        });

    }

    public static boolean isValidEmailAddress(String emailAddress) {
        String emailRegEx;
        Pattern pattern;
        // Regex for a valid email address
        emailRegEx = "^[A-Za-z0-9._%+\\-]+@[A-Za-z0-9.\\-]+\\.[A-Za-z]{2,4}$";
        // Compare the regex with the email address
        pattern = Pattern.compile(emailRegEx);
        Matcher matcher = pattern.matcher(emailAddress);
        if (!matcher.find()) {
            return false;
        }
        return true;
    }

}
