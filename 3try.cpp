#include<bits/stdc++.h>
using namespace std;
int arr1[105];
int main(){
int t;
cin>>t;
while(t--){
int ply;
cin>>ply;
fill(arr1,arr1+105,0);
map<int,map< int, int > > mp;
for(int i=0;i<ply;i++){
  int arr[6]={0};
  int cook;
  cin>>cook;
  for(int j=0;j<cook;j++){
    int c;
    cin>>c;
    mp[i][c-1]++;
}

for (int l=0;l<6;l++){
  arr[l]=mp[i][l];
}

sort(arr,arr+6);
int cost =cook;


for(int l=0;l<6;l++){
  if(arr[l]==0)
  continue;
  else if(arr[l]>0){
    if(l==0)
    {
      cost += 4*(arr[l]);
      for(int y=l+1;y<6;y++)
        arr[y]=arr[y]-arr[l];
    }

    if(l==1)
    {
      cost += 2*(arr[l]);
      for(int y=l+1;y<6;y++)
      arr[y]=arr[y]-arr[l];
    }

    if(l==2)
    {
      cost += 1*(arr[l]);
      for(int y=l+1;y<6;y++)
        arr[y]=arr[y]-arr[l];
    }
  }
}
arr1[i]=cost;
}
int flag =-1,ind = -1,max=-1;

for(int k=0;k<ply;k++){
  if(arr1[k]>max){
    max =arr1[k];
    ind = k;
  }else if(arr1[k]==max){
    flag=2;
  }
}

if(flag==2){
  cout<<"tie"<<endl;
}else{
  if(ind==0){
    cout<<"chef"<<endl;
  }else
  {
    cout<<ind+1<<endl;
  }
}

}
return 0;
}
