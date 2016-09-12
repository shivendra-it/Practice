#include<bits/stdc++.h>
using namespace std;
int main(){
int t;
cin>>t;
while(t--){
char s[12346];
scanf("%s",s);
int f =2;
int l = strlen(s);
for (int i=0,j=l-1;i<l/2;i++,j--){
  if(s[i]==s[j] && !(s[i]== '.' && s[j]=='.' )){
    continue;
  }else if(s[i]== '.' || s[j]=='.'){
    if(s[i]== '.' && s[j]!='.'){
      s[i]=s[j];
    }else if(s[i]!= '.' && s[j]=='.'){
      s[j]=s[i];
    }else if(s[i]== '.' && s[j]=='.'){
      s[i]='a';
      s[j]='a';
    }
  }else if(s[i]!=s[j]){
    f=1;
    break;
  }
}
if(l%2==1){
  int p =l/2;
  if(s[p]== '.')
  s[p]='a';
}

if (f==1){
  cout<<"-1"<<endl;
}else{
  cout<<s<<endl;
}
}
return 0;
}
