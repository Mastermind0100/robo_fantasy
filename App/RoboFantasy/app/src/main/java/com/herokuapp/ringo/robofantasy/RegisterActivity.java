package com.herokuapp.ringo.robofantasy;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.text.TextUtils;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;

import com.google.android.material.textfield.TextInputEditText;

import static com.herokuapp.ringo.robofantasy.MainActivity.isValidEmailAddress;

public class RegisterActivity extends AppCompatActivity {

    private Button btn1;
    private TextInputEditText email;
    private TextInputEditText pass;
    private TextInputEditText namee;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_register);

        namee =  findViewById(R.id.name);
        email = findViewById(R.id.emailid);
        pass = findViewById(R.id.passcode);
        btn1  = findViewById(R.id.clickReg);
        btn1.setOnClickListener(new View.OnClickListener() {
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
                else

                    startActivity(new Intent(RegisterActivity.this, MainActivity.class));
            }
        });
    }
}
