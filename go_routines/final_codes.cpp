//********  Tree  *********

#include <iostream>

using namespace std;

struct TreeNode
{
    int value;
    TreeNode *left, *right;
    TreeNode(int v, TreeNode *l = NULL, TreeNode *r = NULL) : value(v), left(l), right(r) {}
};

int height(TreeNode *root) {
    if (!root) return 0;
    return 1 + max(height(root->left), height(root->right));
}

bool isBalanced(TreeNode *root) {
    if (!root) return true;
    return abs(height(root->left) - height(root->right)) <= 1;
}

int main() {
    TreeNode *root = new TreeNode(0);
    root->left = new TreeNode(1);
    root->right = new TreeNode(2);
    root->left->left = new TreeNode(3);
    root->left->left->left = new TreeNode(4);
    cout << isBalanced(root) << endl;
    cout << height(root) << endl;
    return 0;
}



//********  MAP  *********


#include <iostream>
#include <map>

using namespace std;

int main() {
	map<int, string> a;
	map<int, string>::iterator it;
	a[0] = "This ";
	a[1] = "is";
	a[2] = "funny!\n";
	it = a.find(1);
	if (it != a.end()) {
		cout << it->second << "\n";
	}
	else cout << "1 doesn't exist\n";
	a.erase(it);
	it = a.find(1);
	if (it == a.end()) cout << "1 has been erased!\n";
	a.insert(pair<int, string>(1, "is"));
	it = a.find(1);
	if (it != a.end()) cout << "1 has been inserted back\n";
	return 0;
}

//********  MAP  *********
